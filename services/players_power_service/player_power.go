package players_service

import (
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/utils/errors"
)

var (
	PlayerPowerServices playerPowerServiceInterface = &playersPowerService{}
)

type playersPowerService struct{}

type playerPowerServiceInterface interface {
	PracticePlayer(id int64, paramKey string) *errors.RestErr
	SetDefaultPowers(id int64) *errors.RestErr
}

func (service *playersPowerService) PracticePlayer(id int64, paramKey string) *errors.RestErr {
	dao := domain.PlayerPower{}
	return dao.PlayerPractice(id, paramKey)
}

func (service *playersPowerService) SetDefaultPowers(id int64) *errors.RestErr {
	dao := domain.PlayerPower{}
	return dao.SetDefaultPowers(id)
}
