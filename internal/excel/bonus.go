package excel

import (
	"bytes"
	"code.citik.ru/back/report-action/internal"
	"fmt"
	"github.com/xuri/excelize/v2"
)

const (
	bonusSheetName = "Sheet1"
)

type BonusFileGenerator struct {
}

func NewBonusFileGenerator() *BonusFileGenerator {
	return &BonusFileGenerator{}
}

// GenerateGeneral генерирует файл отчета по общей информации о бонусах
func (b *BonusFileGenerator) GenerateGeneral(bonuses []*internal.BonusGeneral) ([]byte, error) {
	file := excelize.NewFile()
	file.NewSheet(bonusSheetName)

	headers := &[]interface{}{
		"NavActionNumber",
		"Bonus",
		"CampaignStartDate",
		"CampaignFinishDate",
		"CountClients",
		"CountClientsWithActiveCard",
		"CountClientsSendActivation",
		"CountClientsSuccessActivation",
		"ActivationPercent",
	}
	err := file.SetSheetRow(bonusSheetName, "A1", headers)
	if err != nil {
		return nil, fmt.Errorf("can't set header row: %w", err)
	}

	dataRows := 0
	for _, bonus := range bonuses {
		data := &[]interface{}{
			bonus.NavActionNumber,
			bonus.Bonus,
			bonus.CampaignStartDate,
			bonus.CampaignFinishDate,
			bonus.CountClients,
			bonus.CountClientsWithActiveCard,
			bonus.CountClientsSendActivation,
			bonus.CountClientsSuccessActivation,
			bonus.ActivationPercent,
		}

		err = file.SetSheetRow(bonusSheetName, "A2", data)
		if err != nil {
			return nil, fmt.Errorf("can't set row: %w", err)
		}

		dataRows++
	}
	if dataRows == 0 {
		return nil, fmt.Errorf("no data to write")
	}

	var buf bytes.Buffer
	err = file.Write(&buf)
	if err != nil {
		return nil, fmt.Errorf("can't write to buffer: %w", err)
	}

	return buf.Bytes(), nil
}

// GenerateDetailed генерирует файл отчета c детальной информацией о бонусах
func (b *BonusFileGenerator) GenerateDetailed(ch <-chan *internal.BonusDetailed) ([]byte, error) {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter(bonusSheetName)
	if err != nil {
		return nil, fmt.Errorf("can't create stream writer: %w", err)
	}

	headers := map[string]string{
		"A1": "NavActionNumber",
		"B1": "Bonus",
		"C1": "CampaignStartDate",
		"D1": "CampaignFinishDate",
		"E1": "SendForActivationDate",
		"F1": "ActivationDate",
		"G1": "NavOperationNum",
		"H1": "NavClientNum",
		"I1": "LoyaltyCard",
		"J1": "BonusAmountOnBalance",
	}
	for k, v := range headers {
		err = streamWriter.SetRow(k, []interface{}{excelize.Cell{Value: v}}, excelize.RowOpts{Height: 40, Hidden: false})
		if err != nil {
			return nil, fmt.Errorf("can't set header row: %w", err)
		}
	}
	dataRows := 0
	for bonus := range ch {
		row := make([]interface{}, len(headers))
		row[0] = bonus.NavActionNumber
		row[1] = bonus.Bonus
		row[2] = bonus.CampaignStartDate
		row[3] = bonus.CampaignFinishDate
		row[4] = bonus.SendForActivationDate
		row[5] = bonus.ActivationDate
		row[6] = bonus.NavOperationNum
		row[7] = bonus.NavClientNum
		row[8] = bonus.LoyaltyCard
		row[9] = bonus.BonusAmountOnBalance

		cell, _ := excelize.CoordinatesToCellName(1, dataRows+2)
		err = streamWriter.SetRow(cell, row)
		if err != nil {
			return nil, fmt.Errorf("can't set row: %w", err)
		}
		dataRows++
	}
	if dataRows == 0 {
		return nil, fmt.Errorf("no data to write")
	}

	err = streamWriter.Flush()
	if err != nil {
		return nil, fmt.Errorf("can't flush data stream: %w", err)
	}

	var buf bytes.Buffer
	err = file.Write(&buf)
	if err != nil {
		return nil, fmt.Errorf("can't write to buffer: %w", err)
	}

	return buf.Bytes(), nil
}
