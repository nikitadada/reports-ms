package internal

import "code.citik.ru/back/report-action/internal/distributed_work"

const (
	BonusTaskType distributed_work.TaskType = "bonus"
)

type Payload struct {
	ReportId        string
	ActionNumber    string
	ActionStartTime int64
	ReportType      ReportType
}
