package mssql

import (
	"code.citik.ru/back/report-action/internal"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestStorage_FindDetailed(t *testing.T) {
	ctx := context.Background()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	rows := []string{
		"nav_action_no",
		"bonus",
		"campaign_start_date",
		"campaign_finish_date",
		"send_for_activation_date_time",
		"activation_date_time",
		"nav_operation_no",
		"nav_client_no",
		"loyalty_card",
		"bonus_amount_on_balance",
	}
	mock.ExpectQuery("Reports.get_promo_bonuses_activation_details").
		WillReturnRows(sqlmock.NewRows(rows).AddRow(
			"action num3",
			100,
			time.Unix(1624320000, 0),
			time.Unix(1624320000, 0),
			time.Unix(1624320000, 0),
			time.Unix(1624320000, 0),
			"test3",
			"test3",
			"test3",
			50,
		))

	mockLogger := zap.NewNop()

	r := NewBonusReportStorage(db, mockLogger)

	reports := make([]internal.BonusDetailed, 0)

	ch, err := r.FindDetailed(ctx, "action num3", time.Unix(1624320000, 0))

	for report := range ch {
		reports = append(reports, *report)
	}

	assert.NoError(t, err)
	assert.Equal(t, []internal.BonusDetailed{
		{
			"action num3",
			100,
			time.Unix(1624320000, 0),
			time.Unix(1624320000, 0),
			time.Unix(1624320000, 0),
			time.Unix(1624320000, 0),
			"test3",
			"test3",
			"test3",
			50,
		},
	}, reports)
}

func TestRepository_FindDetailedError(t *testing.T) {
	ctx := context.Background()
	mockLogger := zap.NewNop()

	mockDB, mockSql, _ := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mockSql.ExpectQuery("Reports.get_promo_bonuses_activation_details").WillReturnError(errors.New("test"))

	r := &BonusStorage{
		db:     sqlxDB,
		logger: mockLogger,
	}
	ch, err := r.FindDetailed(ctx, "action num", time.Unix(1624320000, 0))

	assert.EqualError(t, err, "database error, see log for more info")
	assert.Nil(t, ch)
}

func TestStorage_FindGeneral(t *testing.T) {
	ctx := context.Background()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	rows := []string{
		"nav_action_no",
		"bonus",
		"count_clients",
		"campaign_start_date",
		"campaign_finish_date",
		"count_clients_with_active_card",
		"count_clients_send_for_bonus_activation",
		"count_clients_succeded_bonus_activation",
		"activation_percent",
	}
	mock.ExpectQuery("Reports.get_promo_bonuses_count").
		WillReturnRows(sqlmock.NewRows(rows).AddRow(
			"action num", 100, 10, time.Unix(1624320000, 0), time.Unix(1624320000, 0), 5, 5, 5, 50),
		)

	mockLogger := zap.NewNop()
	storage := NewBonusReportStorage(db, mockLogger)

	reports, err := storage.FindGeneral(ctx, "action num3", time.Unix(1624320000, 0))

	for _, report := range reports {
		reports = append(reports, report)
	}

	assert.NoError(t, err)
	assert.Equal(t, &internal.BonusGeneral{
		NavActionNumber:               "action num",
		Bonus:                         100,
		CountClients:                  10,
		CampaignStartDate:             time.Unix(1624320000, 0),
		CampaignFinishDate:            time.Unix(1624320000, 0),
		CountClientsWithActiveCard:    5,
		CountClientsSendActivation:    5,
		CountClientsSuccessActivation: 5,
		ActivationPercent:             "50",
	}, reports[0])
}

func TestRepository_FindGeneralError(t *testing.T) {
	ctx := context.Background()
	mockLogger := zap.NewNop()

	mockDB, mockSql, _ := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mockSql.ExpectQuery("Reports.get_promo_bonuses_count").WillReturnError(errors.New("test"))

	r := &BonusStorage{
		db:     sqlxDB,
		logger: mockLogger,
	}
	reports, err := r.FindGeneral(ctx, "action num", time.Unix(1624320000, 0))

	assert.EqualError(t, err, "database error, see log for more info")
	assert.Nil(t, reports)
}

func TestStorage_FindGeneral_ErrorScan(t *testing.T) {
	ctx := context.Background()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	rows := []string{
		"nav_action_no",
		"bonus",
		"count_clients",
		"campaign_start_date",
		"campaign_finish_date",
		"count_clients_with_active_card",
		"count_clients_send_for_bonus_activation",
		"count_clients_succeded_bonus_activation",
		"activation_percent",
	}
	mock.ExpectQuery("Reports.get_promo_bonuses_count").
		WillReturnRows(sqlmock.NewRows(rows).AddRow(
			"action num", "invalid", 10, time.Unix(1624320000, 0), time.Unix(1624320000, 0), 5, 5, 5, 50),
		)

	mockLogger := zap.NewNop()
	storage := NewBonusReportStorage(db, mockLogger)
	reports, err := storage.FindGeneral(ctx, "action num", time.Unix(1624320000, 0))

	assert.NoError(t, err)
	assert.Nil(t, reports)
}

func TestStorage_FindDetailed_ErrorScan(t *testing.T) {
	ctx := context.Background()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	rows := []string{
		"nav_action_no",
		"bonus",
		"campaign_start_date",
		"campaign_finish_date",
		"send_for_activation_date_time",
		"activation_date_time",
		"nav_operation_no",
		"nav_client_no",
		"loyalty_card",
		"bonus_amount_on_balance",
	}
	mock.ExpectQuery("Reports.get_promo_bonuses_activation_details").
		WillReturnRows(sqlmock.NewRows(rows).AddRow(
			"action num3",
			100,
			"invalid",
			time.Unix(1624320000, 0),
			time.Unix(1624320000, 0),
			time.Unix(1624320000, 0),
			"test3",
			"test3",
			"test3",
			50,
		))

	mockLogger := zap.NewNop()

	storage := NewBonusReportStorage(db, mockLogger)
	reports := make([]internal.BonusDetailed, 0)
	ch, err := storage.FindDetailed(ctx, "action num3", time.Unix(1624320000, 0))

	for report := range ch {
		reports = append(reports, *report)
	}

	assert.NoError(t, err)
	assert.Equal(t, reports, []internal.BonusDetailed{})
}
