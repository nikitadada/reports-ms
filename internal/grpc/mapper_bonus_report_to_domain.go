package grpc

import (
	"code.citik.ru/back/report-action/internal"
	bonusv1 "code.citik.ru/back/report-action/internal/grpc/gen/citilink/reportaction/bonus/v1"
	"errors"
)

// bonusReportToDomainMapper маппер отчетов по бонусам из protobuf в доменные структуры
type bonusReportToDomainMapper struct{}

// NewBonusReportToDomainMapper создает новый мапер в доменные структуры.
func NewBonusReportToDomainMapper() *bonusReportToDomainMapper {
	return &bonusReportToDomainMapper{}
}

func (m *bonusReportToDomainMapper) mapType(pbType bonusv1.Type) (internal.ReportType, error) {
	domainType, ok := toDomainEnumMaps.bonusReportTypes[pbType]
	if !ok {
		return 0, errors.New("type is invalid")
	}

	return domainType, nil
}
