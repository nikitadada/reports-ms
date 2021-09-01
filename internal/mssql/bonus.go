package mssql

import (
	"code.citik.ru/back/report-action/internal"
	"code.citik.ru/gobase/database"
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"time"
)

const (
	procGeneral  = "Reports.get_promo_bonuses_count"
	procDetailed = "Reports.get_promo_bonuses_activation_details"
)

// BonusStorage хранилище данных о бонусах акций
type BonusStorage struct {
	db     database.DB
	logger *zap.Logger
}

// NewBonusReportStorage создает новое хранилище бонусов акций
func NewBonusReportStorage(db database.DB, logger *zap.Logger) *BonusStorage {
	return &BonusStorage{db: db, logger: logger}
}

// Ряд выборки с общими данными по бонусам
type bonusGeneralRow struct {
	ActionNumber                  string        `db:"nav_action_no"`
	Bonus                         int           `db:"bonus"`
	CountClients                  int           `db:"count_clients"`
	CampaignStartDate             time.Time     `db:"campaign_start_date"`
	CampaignFinishDate            time.Time     `db:"campaign_finish_date"`
	CountClientsWithActiveCard    int           `db:"count_clients_with_active_card"`
	CountClientsSendActivation    int           `db:"count_clients_send_for_bonus_activation"`
	CountClientsSuccessActivation sql.NullInt32 `db:"count_clients_succeded_bonus_activation"`
	ActivationPercent             string        `db:"activation_percent"`
}

// Ряд выборки с детализированными данными по бонусам
type bonusDetailedRow struct {
	ActionNumber          string        `db:"nav_action_no"`
	Bonus                 int           `db:"bonus"`
	CampaignStartDate     time.Time     `db:"campaign_start_date"`
	CampaignFinishDate    time.Time     `db:"campaign_finish_date"`
	SendForActivationDate time.Time     `db:"send_for_activation_date_time"`
	ActivationDate        sql.NullTime  `db:"activation_date_time"`
	NavOperationNum       string        `db:"nav_operation_no"`
	NavClientNum          string        `db:"nav_client_no"`
	LoyaltyCard           string        `db:"loyalty_card"`
	BonusAmountOnBalance  sql.NullInt32 `db:"bonus_amount_on_balance"`
}

func (b *BonusStorage) FindDetailed(
	ctx context.Context,
	actionNumber string,
	actionStartTime time.Time,
) (<-chan *internal.BonusDetailed, error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Second*35)

	rows, err := b.db.QueryxContext(ctxDb, procDetailed,
		sql.Named("nav_action_no", actionNumber),
		sql.Named("campaign_date_start", actionStartTime),
	)
	if err != nil {
		cancel()
		b.logger.Error("cannot bind or exec named query", zap.Error(err))

		return nil, errors.New("database error, see log for more info")
	}

	ch := make(chan *internal.BonusDetailed, 2)

	go func() {
		defer cancel()
		defer rows.Close()
		defer close(ch)

		for rows.Next() {
			select {
			case <-ctx.Done():
				return
			default:
			}

			raw := &bonusDetailedRow{}
			err = rows.StructScan(&raw)
			if err != nil {
				b.logger.Error("cannot scan row", zap.Error(err))
				continue
			}

			ch <- &internal.BonusDetailed{
				NavActionNumber:       raw.ActionNumber,
				Bonus:                 raw.Bonus,
				CampaignStartDate:     raw.CampaignStartDate,
				CampaignFinishDate:    raw.CampaignFinishDate,
				SendForActivationDate: raw.SendForActivationDate,
				ActivationDate:        raw.ActivationDate.Time,
				NavOperationNum:       raw.NavOperationNum,
				NavClientNum:          raw.NavClientNum,
				LoyaltyCard:           raw.LoyaltyCard,
				BonusAmountOnBalance:  int(raw.BonusAmountOnBalance.Int32),
			}
		}
	}()

	return ch, nil
}

func (b *BonusStorage) FindGeneral(
	ctx context.Context,
	actionNumber string,
	actionStartTime time.Time,
) ([]*internal.BonusGeneral, error) {
	ctxDb, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	rows, err := b.db.QueryxContext(ctxDb, procGeneral,
		sql.Named("nav_action_no", actionNumber),
		sql.Named("campaign_date_start", actionStartTime),
	)
	if err != nil {
		b.logger.Error("cannot bind or exec named query", zap.Error(err))

		return nil, errors.New("database error, see log for more info")
	}
	defer rows.Close()

	var bonuses []*internal.BonusGeneral
	for rows.Next() {
		raw := &bonusGeneralRow{}
		err = rows.StructScan(&raw)
		if err != nil {
			b.logger.Error("cannot scan row", zap.Error(err))
			continue
		}

		bonuses = append(bonuses, &internal.BonusGeneral{
			NavActionNumber:               raw.ActionNumber,
			Bonus:                         raw.Bonus,
			CountClients:                  raw.CountClients,
			CampaignStartDate:             raw.CampaignStartDate,
			CampaignFinishDate:            raw.CampaignFinishDate,
			CountClientsWithActiveCard:    raw.CountClientsWithActiveCard,
			CountClientsSendActivation:    raw.CountClientsSendActivation,
			CountClientsSuccessActivation: int(raw.CountClientsSuccessActivation.Int32),
			ActivationPercent:             raw.ActivationPercent,
		})
	}

	return bonuses, nil
}
