package players_service

import (
	"fmt"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"github.com/kasrashrz/Affogato/utils/errors"
	"github.com/kasrashrz/Affogato/utils/player_utils"
)

var (
	PlayerServices playerServiceInterface = &playersService{}
)

type playersService struct{}

type playerServiceInterface interface {
	GetAll() ([]domain.Player, *errors.RestErr)
	GetOne(id int64) (*domain.Player, *errors.RestErr)
	Create(player domain.Player) *errors.RestErr
	GetPlayersWithoutTeam() ([]domain.Player, *errors.RestErr)
	GenerateRandom(quantity int) *errors.RestErr
	GoalIncrement(id, goal int64) *errors.RestErr
	BodyBuilder(id int64, energy int) *errors.RestErr
	SetDefaultGoals(id int64) *errors.RestErr
	IncreaseAge(id int64) *errors.RestErr
	RemoveOlds() *errors.RestErr
	ChangeStatus(id int64, statusId int64) *errors.RestErr
	Delete(id int64) *errors.RestErr
	PlayerAvg(id int64) (int64, *errors.RestErr)
	PowerList(id int64) ([]domain.PlayerPower, *errors.RestErr)
	RemovePlayerFromTeam(id int64) *errors.RestErr
	SearchP2PWithFilter(filters []dto_only.PlayerFilters) ([]map[string]interface{}, *errors.RestErr)
	SearchWithFilter(filters []dto_only.PlayerFilters) ([]map[string]interface{}, *errors.RestErr)
	AddToBid(bid domain.Bid) *errors.RestErr
	GetAllBids() ([]domain.Bid, *errors.RestErr)
	GetOneBid(bidId int64) (*domain.Bid, *errors.RestErr)
	IncreaseBid(bidId, buyerId, price int64) *errors.RestErr
}

func (service *playersService) GetAll() ([]domain.Player, *errors.RestErr) {
	dao := domain.Player{}
	return dao.GetAll()
}

func (service *playersService) GetOne(id int64) (*domain.Player, *errors.RestErr) {
	dao := domain.Player{}
	result, err := dao.GetOne(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *playersService) Create(player domain.Player) *errors.RestErr {

	if player_utils.PowerCheck(player.Power) {
		return errors.BadRequestError("parameters value must be at least 25")
	}
	return player.Create()
}

func (service *playersService) GetPlayersWithoutTeam() ([]domain.Player, *errors.RestErr) {
	dao := domain.Player{}
	return dao.GetPlayersWithoutTeam()
}

func (service *playersService) GenerateRandom(quantity int) *errors.RestErr {
	dao := domain.Player{}
	return dao.GenerateRandom(quantity)
}

func (service *playersService) GoalIncrement(id, goal int64) *errors.RestErr {
	dao := domain.Player{}
	return dao.GoalIncrement(id, goal)
}

func (service *playersService) BodyBuilder(id int64, energy int) *errors.RestErr {
	dao := domain.Player{}
	return dao.BodyBuilder(id, energy)
}

func (service *playersService) SetDefaultGoals(id int64) *errors.RestErr {
	dao := domain.Player{}
	return dao.SetDefaultGoals(id)
}

func (service *playersService) IncreaseAge(id int64) *errors.RestErr {
	dao := domain.Player{}
	return dao.IncreaseAge(id)
}

func (service *playersService) RemoveOlds() *errors.RestErr {
	dao := domain.Player{}
	return dao.RemoveOlds()
}

func (service *playersService) ChangeStatus(id int64, statusId int64) *errors.RestErr {
	dao := domain.Player{}
	return dao.ChangeStatus(id, statusId)
}

func (service *playersService) Delete(id int64) *errors.RestErr {
	dao := domain.Player{}
	return dao.Delete(id)
}

func (service *playersService) PlayerAvg(id int64) (int64, *errors.RestErr) {
	dao := domain.Player{}
	return dao.PlayerAvg(id)
}

func (service *playersService) PowerList(id int64) ([]domain.PlayerPower, *errors.RestErr) {
	dao := domain.Player{}
	return dao.PowerList(id)
}

func (service *playersService) RemovePlayerFromTeam(id int64) *errors.RestErr {
	dao := domain.Player{}
	return dao.RemovePlayerFromTeam(id)
}

func (service *playersService) SearchP2PWithFilter(filters []dto_only.PlayerFilters) ([]map[string]interface{}, *errors.RestErr) {
	dao := domain.Player{}
	queryString := ""
	for i, filter := range filters {
		if len(filters)-i != 1 {
			queryString += filter.Key + " " + filter.Operation + " " + filter.Value + " AND "
		}
		if len(filters)-i == 1 {
			queryString += filter.Key + " " + filter.Operation + " " + filter.Value + " "
		}
	}
	fmt.Println(queryString)

	return dao.SearchP2PWithFilter(queryString)
}

func (service *playersService) SearchWithFilter(filters []dto_only.PlayerFilters) ([]map[string]interface{}, *errors.RestErr) {
	dao := domain.Player{}
	queryString := ""
	for i, filter := range filters {
		if len(filters)-i != 1 {
			queryString += filter.Key + " " + filter.Operation + " " + filter.Value + " AND "
		}
		if len(filters)-i == 1 {
			queryString += filter.Key + " " + filter.Operation + " " + filter.Value + " "
		}
	}
	fmt.Println(queryString)

	return dao.SearchWithFilter(queryString)
}

func (service *playersService) AddToBid(bid domain.Bid) *errors.RestErr {
	dao := domain.Player{}
	return dao.AddToBid(bid)
}

func (service *playersService) GetAllBids() ([]domain.Bid, *errors.RestErr) {
	dao := domain.Player{}
	return dao.GetAllBids()
}

func (service *playersService) GetOneBid(bidId int64) (*domain.Bid, *errors.RestErr) {
	dao := domain.Player{}
	return dao.GetOneBid(bidId)
}

func (service *playersService) IncreaseBid(bidId, buyerId, price int64) *errors.RestErr {
	dao := domain.Player{}
	bid, _ := dao.GetOneBid(bidId)
	if bid.Price > price {
		return errors.BadRequestError("invalid price, not enough coins")
	}
	return dao.IncreaseBid(bidId, buyerId, price)
}
