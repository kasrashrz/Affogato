package market_service

import (
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/utils/errors"
	"gorm.io/gorm"
)

var (
	MarketServices marketServiceInterface = &marketsService{}
)

type marketsService struct{}

type marketServiceInterface interface {
	AddPlayer(market domain.Market) *errors.RestErr
	BuyPlayer(uid int64, marketId int64) *errors.RestErr
	AllMarketChoices(start, count int64) ([]map[string]interface{}, *errors.RestErr)
}

func (service *marketsService) AddPlayer(market domain.Market) *errors.RestErr {
	dao := domain.Market{
		Model:         gorm.Model{},
		Player:        domain.Player{},
		PlayerId:      market.PlayerId,
		PlayerPrice:   market.PlayerPrice,
		AutoGenerated: false,
		IsSold:        market.IsSold,
	}
	return dao.AddPlayer()
}

func (service *marketsService) BuyPlayer(uid int64, marketId int64) *errors.RestErr {
	dao := domain.Market{}
	return dao.BuyPlayer(uid, marketId)
}

func (service *marketsService) AllMarketChoices(start, count int64) ([]map[string]interface{}, *errors.RestErr) {
	dao := domain.Market{}
	return dao.AllMarketChoices(start, count)
}
