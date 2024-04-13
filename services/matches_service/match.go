package matches_service

import (
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/utils/errors"
)

var (
	MatchServices matchServiceInterface = &matchesService{}
)

type matchesService struct{}

type matchServiceInterface interface {
	Create(match domain.Match) *errors.RestErr
	MatchNotification(id int64) ([]domain.MatchNotification, *errors.RestErr)
	UpdateResult(match domain.Match) *errors.RestErr
	UsersMatches(id int64) ([]domain.Match, *errors.RestErr)
	SingleMatch(id int64) (*domain.Match, *errors.RestErr)
	FriendlyMatch(match domain.Match) *errors.RestErr
	JoinFriendlyMatch(uid int64) *errors.RestErr
	GetFriendlyMatchesInvites(uid int64) ([]domain.MatchNotification, *errors.RestErr)
}

func (matchService *matchesService) Create(match domain.Match) *errors.RestErr {

	return match.Create()
}

func (matchService *matchesService) MatchNotification(id int64) ([]domain.MatchNotification, *errors.RestErr) {
	dao := domain.Match{}
	result, err := dao.MatchNotification(id)
	return result, err
}

func (matchService *matchesService) UpdateResult(match domain.Match) *errors.RestErr {
	return match.Update()
}

func (matchService *matchesService) UsersMatches(id int64) ([]domain.Match, *errors.RestErr) {
	dao := domain.Match{}
	return dao.UsersMatches(id)
}

func (matchService *matchesService) SingleMatch(id int64) (*domain.Match, *errors.RestErr) {
	dao := domain.Match{}
	return dao.SingleMatch(id)
}

func (matchService *matchesService) FriendlyMatch(match domain.Match) *errors.RestErr {
	dao := domain.Match{}
	return dao.FriendlyMatch(match)
}

func (matchService *matchesService) JoinFriendlyMatch(uid int64) *errors.RestErr {
	dao := domain.Match{}
	return dao.JoinFriendlyMatch(uid)
}

func (matchService *matchesService) GetFriendlyMatchesInvites(uid int64) ([]domain.MatchNotification, *errors.RestErr) {
	dao := domain.Match{}
	return dao.GetFriendlyMatchesInvites(uid)
}
