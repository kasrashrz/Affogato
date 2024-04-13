package domain

import (
	checkErr "errors"
	"fmt"
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"github.com/kasrashrz/Affogato/logger"
	"github.com/kasrashrz/Affogato/utils/errors"
	"gorm.io/gorm"
	"math/rand"
	"strings"
	"time"
)

func (player *Player) GetAll() ([]Player, *errors.RestErr) {
	var players []Player

	if err := db.Preload("Power").
		Preload("Status").
		Preload("Post").
		Table("players").
		Scan(&players).Limit(200).Error; err != nil {
		logger.Error("read all players", err)
		return nil, errors.InternalServerError("something went wrong")
	}

	return players, nil
}

func (player *Player) GetOne(id int64) (*Player, *errors.RestErr) {
	var outputPlayer Player

	if err := db.Preload("Power").Preload("Status").Preload("Post").
		Where("id = ?", id).First(&outputPlayer).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFoundError("player not found")
		}
		return nil, errors.InternalServerError("something went wrong")
	}

	return &outputPlayer, nil
}

func (player *Player) Create() *errors.RestErr {

	if err := db.Save(&player).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (player *Player) GetPlayersWithoutTeam() ([]Player, *errors.RestErr) {
	var players []Player

	if err := db.Table("players").Where("team_id IS NULL").Scan(&players).Limit(100); err.Error != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return players, nil
}

func (player *Player) GenerateRandom(quantity int) *errors.RestErr {

	var FirstNames = []string{"ali", "asghar", "jafar", "hosein"}
	var LastNames = []string{"jafari", "sadri", "shirazi", "hoseini"}

	for i := 0; i < quantity; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 2
		insertPlayer := Player{
			Model: gorm.Model{},
			Name:  FirstNames[rand.Intn(max-min+1)] + " " + LastNames[rand.Intn(max-min+1)],
			Post: Post{
				Model: gorm.Model{
					ID: uint(rand.Intn(4-1+1) + 1),
				},
			},
			Status: dto_only.Status{
				Model: gorm.Model{
					ID: uint(rand.Intn(4-1+1) + 0),
				},
			},
			Age:    rand.Intn(26-18+1) + 18,
			Skill:  rand.Intn(3 - 1 + 1),
			Salary: int64(rand.Intn(2000-100+100) * 5),
			Cheer:  0,
			Goal:   rand.Intn(120-0+1) + 0,
			Price:  int64(rand.Intn(100000-10000+1) + 10000),
		}

		if err := db.Create(&insertPlayer).Error; err != nil {
			if strings.Contains(err.Error(), "Error 1452") {
				return errors.BadRequestError("child row does not exist")
			}
			return errors.InternalServerError("something went wrong")
		}

		playerPower := PlayerPower{
			Model:    gorm.Model{},
			PlayerID: insertPlayer.ID,
		}
		if err := db.Create(&playerPower).Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}

	}

	return nil
}

func (player *Player) GoalIncrement(id, goals int64) *errors.RestErr {

	var findPlayer Player
	err := db.Table("players").Where("id = ?", id).First(&findPlayer).Error
	if checkErr.Is(err, gorm.ErrRecordNotFound) {
		return errors.BadRequestError("player does not exist")
	}

	if err := db.Table("players").
		Where("id = ?", id).
		Update("goal", gorm.Expr("goal + ?", goals)).Error; err != nil {
		logger.Error("something went wrong when updating player", err)
		return errors.InternalServerError("something went wrong")
	}

	return nil

}

func (player *Player) BodyBuilder(id int64, energy int) *errors.RestErr {

	var findPlayer Player
	err := db.Table("players").Where("id = ?", id).First(&findPlayer).Error
	if checkErr.Is(err, gorm.ErrRecordNotFound) {
		return errors.BadRequestError("player does not exist")
	}

	if findPlayer.Energy+energy > 100 {
		return errors.BadRequestError("too much energy")
	}

	if err := db.Table("players").
		Where("id = ?", id).
		Update("energy", gorm.Expr("energy + ?", energy)).Error; err != nil {
		logger.Error("something went wrong when updating player", err)
		return errors.InternalServerError("something went wrong")
	}

	return nil

}

func (player *Player) SetDefaultGoals(id int64) *errors.RestErr {

	if err := db.Table("players").Where("id = ?", id).Update("goal", 0).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (player *Player) IncreaseAge(id int64) *errors.RestErr {

	if err := db.Table("players").
		Where("id = ?", id).
		Update("age", gorm.Expr("age + ?", 1)).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (player *Player) RemoveOlds() *errors.RestErr {

	if err := db.Exec("delete players, pp from players join player_powers pp on players.id = pp.player_id where players.age > 26").Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (player *Player) ChangeStatus(id int64, statusId int64) *errors.RestErr {

	if err := db.Table("players").Where("id = ?", id).Update("status_id", statusId).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (player *Player) Delete(id int64) *errors.RestErr {

	if err := db.Exec("delete players, pp from players join player_powers pp on play"+
		"ers.id = pp.player_id where players.id = ?", id).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("user does not exist")
		}
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (player *Player) PlayerAvg(id int64) (avg int64, err *errors.RestErr) {

	averageQuery := "AVG(CAST(shoot + pass + strength + endurance + dribble + control + head + pp.goal + pp.power + tackle AS float))"
	if err := db.Select(averageQuery).Table("players").
		Joins("inner join player_powers pp on players.id = pp.player_id").
		Where("players.id = ?", id).
		Scan(&avg).Error; err != nil {
		if strings.ContainsAny(err.Error(), "NULL to int64") {
			return 0, errors.NotFoundError("player does not exist")
		}
		return 0, errors.InternalServerError("something went wrong")
	}
	return avg, nil
}

func (player *Player) PowerList(id int64) ([]PlayerPower, *errors.RestErr) {

	var powerList []PlayerPower
	if err := db.Select("*").Table("players").
		Joins("inner join player_powers pp on players.id = pp.player_id").
		Where("players.id = ?", id).
		Scan(&powerList).Error; err != nil {
		if strings.ContainsAny(err.Error(), "NULL to int64") {
			return nil, errors.NotFoundError("player does not exist")
		}
		return nil, errors.InternalServerError("something went wrong")
	}
	return powerList, nil
}

func (player *Player) RemovePlayerFromTeam(id int64) *errors.RestErr {
	if err := db.Table("players").
		Where("id = ?", id).
		Updates(map[string]interface{}{"team_id": nil}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

// Search inside market
func (player *Player) SearchP2PWithFilter(queryString string) ([]map[string]interface{}, *errors.RestErr) {
	//var players []PlayerFilter
	output := []map[string]interface{}{}
	if err := db.Raw("select" +
		" bids.id as `bid_id`, bids.player_id, bids.price, p.id, post_id, status_id, name, energy, age, skill, cheer, pp.goal, post_name," +
		" control, power, pass, shoot, dribble, tackle, head, endurance, strength" +
		" from bids" +
		" join players p on bids.player_id = p.id join posts p2 on p.post_id = p2.id" +
		" join player_powers pp on p.id = pp.player_id where bids.is_done = 0 and " + queryString + " limit 30").Scan(&output).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return output, nil
}

func (player *Player) SearchWithFilter(queryString string) ([]map[string]interface{}, *errors.RestErr) {
	//var players []PlayerFilter
	output := []map[string]interface{}{}
	if err := db.Raw("select" +
		" bids.id as `bid_id`, bids.player_id, bids.price, p.id, post_id, status_id, name, energy, age, skill, cheer, pp.goal, post_name," +
		" control, power, pass, shoot, dribble, tackle, head, endurance, strength" +
		" from bids" +
		" join players p on bids.player_id = p.id join posts p2 on p.post_id = p2.id" +
		" join player_powers pp on p.id = pp.player_id where " + queryString + " limit 30").Scan(&output).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return output, nil
}

func (player *Player) AddToBid(bid Bid) *errors.RestErr {
	expTime := time.Now().AddDate(0, 0, 1)
	bid.EXP = expTime.Unix()
	bid.IsDone = false

	var bidExists Bid

	db.Table("bids").Where("player_id = ? and is_done = 0", bid.PlayerID).Scan(&bidExists)

	if bidExists.ID != 0 {
		return errors.BadRequestError("player already exists in bids")
	}

	if err := db.Save(&bid).Error; err != nil {
		if strings.Contains(err.Error(), "Error 1452") {
			return errors.BadRequestError("invalid seller, buyer, or player id")
		}
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (player *Player) GetAllBids() ([]Bid, *errors.RestErr) {
	var bids []Bid
	if err := db.Limit(50).Where("is_done = 0").Find(&bids).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return bids, nil
}

func (player *Player) GetOneBid(bidId int64) (*Bid, *errors.RestErr) {
	var bid Bid
	if err := db.Table("bids").Where("id = ?", bidId).Find(&bid).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return &bid, nil
}

func (player *Player) IncreaseBid(bidId, buyerId, price int64) *errors.RestErr {
	if err := db.Table("bids").Where("id = ?", bidId).
		Updates(map[string]interface{}{"buyer_id": buyerId, "price": price}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (player *Player) BidCron() *errors.RestErr {
	var bids []Bid
	var buyerTeam uint
	var buyerCurrentCoin int64
	var sellerTeam uint

	db.Table("bids").
		Where("unix_timestamp(now()) > bids.exp").
		Scan(&bids)

	for _, bid := range bids {
		fmt.Println(bid.SellerID)
		fmt.Println(bid.BuyerID)
		db.Select("team_id").Table("users").Where("team_id = ?", bid.BuyerID).Scan(&buyerTeam)
		db.Select("team_id").Table("users").Where("team_id = ?", bid.SellerID).Scan(&sellerTeam)
		fmt.Println(sellerTeam)

		db.Select("team_id").Table("users").
			Where("id = ?", bid.BuyerID).
			Scan(&buyerTeam)

		db.Select("team_id").Table("users").
			Where("id = ?", bid.SellerID).
			Scan(&sellerTeam)

		db.Select("coin").Table("users").
			Where("id = ?", bid.BuyerID).
			Scan(&buyerCurrentCoin)

		if err := db.Table("users").Where("id = ?", bid.BuyerID).
			Update("coin", gorm.Expr("coin - ?", bid.Price)).Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}

		if err := db.Table("players").Where("id = ?", bid.PlayerID).
			Updates(map[string]interface{}{"team_id": buyerTeam}).Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}

		if err := db.Table("users").Where("id = ?", bid.SellerID).
			Update("coin", gorm.Expr("coin + ?", bid.Price)).Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}

		if err := db.Table("bids").Where("id = ?", bid.ID).
			Updates(map[string]interface{}{"is_done": 1}).Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}

		if err := db.Table("bids").Where("id = ?", bid.ID).
			Updates(map[string]interface{}{"is_done": 1}).Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}

		sellerPurchase := PaymentHistory{
			Model:           gorm.Model{},
			Team:            Team{},
			TeamId:          int64(sellerTeam),
			PaymentDetail:   PaymentDetail{},
			PaymentDetailId: 13,
			Player:          Player{},
			PlayerId:        int64(bid.PlayerID),
			IsBuy:           false,
		}

		buyerPurchase := PaymentHistory{
			Model:           gorm.Model{},
			PaymentDetail:   PaymentDetail{},
			Team:            Team{},
			TeamId:          int64(buyerTeam),
			PaymentDetailId: 6,
			Player:          Player{},
			PlayerId:        int64(bid.PlayerID),
			IsBuy:           false,
		}

		db.Create(&buyerPurchase)
		db.Create(&sellerPurchase)
		buyerTeam = 0
		sellerTeam = 0
		buyerPurchase = PaymentHistory{}
		sellerPurchase = PaymentHistory{}

	}

	return nil
}
