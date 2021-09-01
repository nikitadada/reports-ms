package internal

import (
	"bytes"
	"code.citik.ru/back/report-action/internal/distributed_work"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"go.uber.org/zap"
	"time"
)

type BonusStorage interface {
	// FindGeneral выбирает обобщенные данные о бонусах из БД
	FindGeneral(
		ctx context.Context,
		actionNumber string,
		ActionStartTime time.Time,
	) ([]*BonusGeneral, error)

	// FindDetailed выбирает все данные о бонусах из БД
	FindDetailed(
		ctx context.Context,
		actionNumber string,
		ActionStartTime time.Time,
	) (<-chan *BonusDetailed, error)
}

// FileGenerator Генератор excel файла с бонусами по акциям
type FileGenerator interface {
	GenerateGeneral(bonuses []*BonusGeneral) ([]byte, error)
	GenerateDetailed(ch <-chan *BonusDetailed) ([]byte, error)
}

type BonusGeneral struct {
	NavActionNumber               string
	Bonus                         int
	CountClients                  int
	CampaignStartDate             time.Time
	CampaignFinishDate            time.Time
	CountClientsWithActiveCard    int
	CountClientsSendActivation    int
	CountClientsSuccessActivation int
	ActivationPercent             string
}

type BonusDetailed struct {
	NavActionNumber       string
	Bonus                 int
	CampaignStartDate     time.Time
	CampaignFinishDate    time.Time
	SendForActivationDate time.Time
	ActivationDate        time.Time
	NavOperationNum       string
	NavClientNum          string
	LoyaltyCard           string
	BonusAmountOnBalance  int
}

// BonusConsumer обработчик генерации отчетов
type BonusConsumer struct {
	storage    BonusStorage
	repository ReportRepository
	logger     *zap.Logger
	appVersion *version.Version
	generator  FileGenerator
	uploader   FileUploader
}

func NewBonusConsumer(
	storage BonusStorage,
	repository ReportRepository,
	logger *zap.Logger,
	appVersion *version.Version,
	generator FileGenerator,
	uploader FileUploader,
) *BonusConsumer {
	return &BonusConsumer{
		storage:    storage,
		repository: repository,
		logger:     logger,
		appVersion: appVersion,
		generator:  generator,
		uploader:   uploader,
	}
}

// CanConsume проверяет, может ли быть выполнена задача данным обработчиком
func (c *BonusConsumer) CanConsume(task distributed_work.Task) bool {
	return task.Type() == BonusTaskType
}

func (c *BonusConsumer) ResolveVersion(v *version.Version) distributed_work.VersionResolverVerdict {
	if v.LessThan(c.appVersion) {
		return distributed_work.VersionResolverVerdictDelete
	} else if v.GreaterThan(c.appVersion) {
		return distributed_work.VersionResolverVerdictReleaseWithDelay
	}

	return distributed_work.VersionResolverVerdictOK
}

// Consume выполняет работу по генерации отчета по бонусам акции
func (c *BonusConsumer) Consume(ctx context.Context, task distributed_work.Task) error {
	c.logger.Debug("consumer is started")
	var payload Payload
	err := json.Unmarshal(task.Payload(), &payload)
	if err != nil {
		return fmt.Errorf("can't unmarshal payload of bonus task: %w", err)
	}

	reportId, err := uuid.Parse(payload.ReportId)
	if err != nil {
		return fmt.Errorf("can't parse reportId to uuid: %w", err)
	}

	setErrorStatus := func() error {
		err := c.repository.Update(
			ctx,
			ReportId(reportId),
			payload.ActionNumber,
			payload.ActionStartTime,
			ReportStatusError,
			payload.ReportType,
		)
		if err != nil {
			return fmt.Errorf("can't update report: %w", err)
		}

		return nil
	}

	var file []byte
	if payload.ReportType == ReportTypeGeneral {
		bonuses, err := c.storage.FindGeneral(ctx, payload.ActionNumber, time.Unix(payload.ActionStartTime, 0))
		if err != nil {
			return fmt.Errorf("failed getting general bonuses: %w", err)
		}

		file, err = c.generator.GenerateGeneral(bonuses)
		if err != nil {
			updateStatusErr := setErrorStatus()
			if updateStatusErr != nil {
				err = fmt.Errorf("%w: can't set error status: %s", err, updateStatusErr)
			}

			c.logger.Error("can't generate general report", zap.Error(err))
			return nil
		}
	} else if payload.ReportType == ReportTypeDetailed {
		ch, err := c.storage.FindDetailed(ctx, payload.ActionNumber, time.Unix(payload.ActionStartTime, 0))
		if err != nil {
			return fmt.Errorf("failed getting detailed bonuses: %w", err)
		}

		file, err = c.generator.GenerateDetailed(ch)
		if err != nil {
			updateStatusErr := setErrorStatus()
			if updateStatusErr != nil {
				err = fmt.Errorf("%w: can't set error status: %s", err, updateStatusErr)
			}

			c.logger.Error("can't generate detailed report", zap.Error(err))
			return nil
		}
	} else {
		c.logger.Error("invalid report type")
		return nil
	}

	uploadStatus := ReportStatusInProcess
	err = c.repository.Update(
		ctx,
		ReportId(reportId),
		payload.ActionNumber,
		payload.ActionStartTime,
		uploadStatus,
		payload.ReportType,
	)
	if err != nil {
		return fmt.Errorf("can't update report in process status: %w", err)
	}

	err = c.uploader.Upload(ctx, bytes.NewReader(file), ReportId(reportId))
	if err != nil {
		updateStatusErr := setErrorStatus()
		if updateStatusErr != nil {
			err = fmt.Errorf("%w: can't set error status: %s", err, updateStatusErr)
		}

		return err
	}

	err = c.repository.Update(
		ctx,
		ReportId(reportId),
		payload.ActionNumber,
		payload.ActionStartTime,
		ReportStatusSuccess,
		payload.ReportType,
	)
	if err != nil {
		c.logger.Error("can't update report success status", zap.Error(err))
		return nil
	}

	c.logger.Debug(fmt.Sprintf("report id: %s", reportId.String()))
	return nil
}
