package leagues_service

import (
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/utils/errors"
)

var (
	LeagueServices leagueServiceInterface = &leaguesService{}
)

type leaguesService struct{}

type leagueServiceInterface interface {
	Create(league domain.League) *errors.RestErr
	TestAPI() *errors.RestErr
	LeaderBoard(uid int64) ([]map[string]interface{}, *errors.RestErr)
	TodayMatches(userId int64) ([]map[string]interface{}, *errors.RestErr)
}

func (service *leaguesService) Create(league domain.League) *errors.RestErr {
	dao := domain.League{}
	return dao.Create()
}

func (service *leaguesService) TestAPI() *errors.RestErr {
	dao := domain.League{}
	return dao.AutoLeagueGenerate()
}

func (service *leaguesService) LeaderBoard(uid int64) ([]map[string]interface{}, *errors.RestErr) {
	dao := domain.League{}
	teams, err := dao.LeaderBoard(uid)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (service *leaguesService) TodayMatches(userId int64) ([]map[string]interface{}, *errors.RestErr) {
	dao := domain.League{}
	todayMatches, err := dao.TodayMatches(userId)
	if err != nil {
		return nil, err
	}
	return todayMatches, nil
}
