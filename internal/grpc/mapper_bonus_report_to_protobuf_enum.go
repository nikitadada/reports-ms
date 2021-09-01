package grpc

import (
	"code.citik.ru/back/report-action/internal"
	bonusv1 "code.citik.ru/back/report-action/internal/grpc/gen/citilink/reportaction/bonus/v1"
)

var toProtobufEnumMaps = struct {
	bonusReportStatuses map[internal.ReportStatus]bonusv1.Status
	bonusReportTypes    map[internal.ReportType]bonusv1.Type
}{
	bonusReportStatuses: map[internal.ReportStatus]bonusv1.Status{
		internal.ReportStatusCreated:   bonusv1.Status_STATUS_CREATED,
		internal.ReportStatusInProcess: bonusv1.Status_STATUS_IN_PROCESS,
		internal.ReportStatusError:     bonusv1.Status_STATUS_ERROR,
		internal.ReportStatusSuccess:   bonusv1.Status_STATUS_SUCCESS,
	},
	bonusReportTypes: map[internal.ReportType]bonusv1.Type{
		internal.ReportTypeGeneral:  bonusv1.Type_TYPE_GENERAL,
		internal.ReportTypeDetailed: bonusv1.Type_TYPE_DETAILED,
	},
}
