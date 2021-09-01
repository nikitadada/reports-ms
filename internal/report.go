package internal

import (
	"context"
	"github.com/google/uuid"
	"time"
)

const (
	// ReportTypeGeneral Тип отчета обобщенный
	ReportTypeGeneral ReportType = 1
	// ReportTypeDetailed Тип отчета детализированный
	ReportTypeDetailed ReportType = 2

	// ReportStatusCreated Запрос на создание отчета создан
	ReportStatusCreated ReportStatus = 1
	// ReportStatusInProcess Отчет в процессе формирования
	ReportStatusInProcess ReportStatus = 2
	// ReportStatusError Формирование отчета завершилось с ошибкой
	ReportStatusError ReportStatus = 3
	// ReportStatusSuccess Отчет успешно сформирован
	ReportStatusSuccess ReportStatus = 4

	// ReportFileExtension Расширение файла отчета
	ReportFileExtension string = "xlsx"
)

// ReportRepository хранилище отчетов
type ReportRepository interface {
	// Filter Ищет отчеты в хранилище
	Filter(ctx context.Context, filter *Filter) ([]*Report, error)
	// Create Создает новый отчет в хранилище
	Create(ctx context.Context, report *Report) error
	// Get Ищет отчет в хранилище по идентификатору
	Get(ctx context.Context, id ReportId) (*Report, error)
	// Update Обновляет отчет в хранилище
	Update(ctx context.Context, id ReportId, actionNumber string, actionStart int64, status ReportStatus, reportType ReportType) error
}

// Filter Структура для фильтрации информации об отчетах
type Filter struct {
	NavActionNumber   string
	CampaignStartDate int64
	ValidTime         time.Time
	Types             []ReportType
	Statuses          []ReportStatus
}

// ReportId ID отчета
type ReportId uuid.UUID

// ReportType тип отчета
type ReportType int

// ReportStatus статус формирования отчета
type ReportStatus int

type Report struct {
	Id                ReportId
	NavActionNumber   string
	CampaignStartDate time.Time
	LastModified      time.Time
	Status            ReportStatus
	Type              ReportType
	FileName          string
}
