package excel

import (
	"code.citik.ru/back/report-action/internal"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestBonusFileGenerator_GenerateGeneral(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fileGenerator := NewBonusFileGenerator()

		buf, err := fileGenerator.GenerateGeneral([]*internal.BonusGeneral{
			{
				NavActionNumber:               "num",
				Bonus:                         100,
				CountClients:                  1,
				CampaignStartDate:             time.Unix(199999999, 0),
				CampaignFinishDate:            time.Unix(299999999, 0),
				CountClientsWithActiveCard:    1,
				CountClientsSendActivation:    1,
				CountClientsSuccessActivation: 1,
				ActivationPercent:             "100",
			},
		})

		assert.Equal(t, reflect.TypeOf(buf).String(), "[]uint8")
		assert.NoError(t, err)
	})

	t.Run("empty data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fileGenerator := NewBonusFileGenerator()

		buf, err := fileGenerator.GenerateGeneral([]*internal.BonusGeneral{})

		assert.Nil(t, buf)
		assert.EqualError(t, err, "no data to write")
	})
}

func TestBonusFileGenerator_GenerateDetailed(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fileGenerator := NewBonusFileGenerator()

		bonusLine := &internal.BonusDetailed{
			NavActionNumber:       "num",
			Bonus:                 100,
			CampaignStartDate:     time.Unix(199999999, 0),
			CampaignFinishDate:    time.Unix(299999999, 0),
			SendForActivationDate: time.Unix(299999999, 0),
			ActivationDate:        time.Unix(299999999, 0),
			NavClientNum:          "client num",
			NavOperationNum:       "op num",
			LoyaltyCard:           "card",
			BonusAmountOnBalance:  100,
		}

		ch := make(chan *internal.BonusDetailed, 2)
		ch <- bonusLine
		close(ch)

		buf, err := fileGenerator.GenerateDetailed(ch)

		assert.Equal(t, reflect.TypeOf(buf).String(), "[]uint8")
		assert.NoError(t, err)
	})
	t.Run("empty data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fileGenerator := NewBonusFileGenerator()

		ch := make(chan *internal.BonusDetailed, 2)
		close(ch)

		buf, err := fileGenerator.GenerateDetailed(ch)

		assert.Nil(t, buf)
		assert.EqualError(t, err, "no data to write")
	})
}
