package grpc

import (
	"code.citik.ru/back/report-action/internal"
	bonusv1 "code.citik.ru/back/report-action/internal/grpc/gen/citilink/reportaction/bonus/v1"
)

var toDomainEnumMaps = struct {
	bonusReportStatuses map[bonusv1.Status]internal.ReportStatus
	bonusReportTypes    map[bonusv1.Type]internal.ReportType
}{
	bonusReportStatuses: map[bonusv1.Status]internal.ReportStatus{
		bonusv1.Status_STATUS_CREATED:    internal.ReportStatusCreated,
		bonusv1.Status_STATUS_IN_PROCESS: internal.ReportStatusInProcess,
		bonusv1.Status_STATUS_ERROR:      internal.ReportStatusError,
		bonusv1.Status_STATUS_SUCCESS:    internal.ReportStatusSuccess,
	},
	bonusReportTypes: map[bonusv1.Type]internal.ReportType{
		bonusv1.Type_TYPE_GENERAL:  internal.ReportTypeGeneral,
		bonusv1.Type_TYPE_DETAILED: internal.ReportTypeDetailed,
	},
}
