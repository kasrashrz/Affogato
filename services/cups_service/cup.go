package cups_service

import (
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/utils/errors"
)

var (
	CupServices cupServiceInterface = &cupsService{}
)

type cupsService struct{}

type cupServiceInterface interface {
	Create(cup domain.Cup, uid int64) *errors.RestErr
	GetAll() ([]domain.Cup, *errors.RestErr)
	GetOne(cupId int64) (*domain.Cup, *errors.RestErr)
	Join(uid, cupId int64) *errors.RestErr
	GetCupMatches(cupId int64) ([]map[string]interface{}, *errors.RestErr)
	UserLeagueAndCupData(uid int64) (map[string]uint, *errors.RestErr)
}

var maxCapacityCondition int

func (service *cupsService) Create(cup domain.Cup, uid int64) *errors.RestErr {
	dao := domain.Cup{}
	if cup.MaxCapacity > 16 {
		return errors.BadRequestError("cup capacity is more than 16")
	}

	if cup.MaxCapacity == 16 {
		maxCapacityCondition = 1
	}
	if cup.MaxCapacity == 8 {
		maxCapacityCondition = 1
	}
	if cup.MaxCapacity == 4 {
		maxCapacityCondition = 1
	}

	if maxCapacityCondition == 0 {
		return errors.BadRequestError("cup capacity must be 4 or 8 or 16")
	}

	return dao.CreateCup(cup, uid)
}

func (service *cupsService) GetAll() ([]domain.Cup, *errors.RestErr) {
	dao := domain.Cup{}
	return dao.GetAllCups()
}

func (service *cupsService) GetOne(cupId int64) (*domain.Cup, *errors.RestErr) {
	dao := domain.Cup{}
	return dao.GetOneCup(cupId)
}

func (service *cupsService) Join(uid, cupId int64) *errors.RestErr {
	dao := domain.Cup{}
	return dao.JoinCup(uid, cupId)
}

func (service *cupsService) GetCupMatches(cupId int64) ([]map[string]interface{}, *errors.RestErr) {
	dao := domain.Cup{}
	return dao.GetCupMatches(cupId)
}

func (service *cupsService) UserLeagueAndCupData(uid int64) (map[string]uint, *errors.RestErr) {
	dao := domain.Cup{}
	return dao.UserLeagueAndCupData(uid)
}
