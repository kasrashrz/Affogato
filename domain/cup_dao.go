package domain

import (
	checkErr "errors"
	"fmt"
	"github.com/kasrashrz/Affogato/utils/errors"
	"gorm.io/gorm"
	"time"
)

func (Entrycup *Cup) CreateCup(cup Cup, uid int64) *errors.RestErr {
	var teamId uint

	if err := db.Select("team_id").Table("users").
		Where("id = ?", uid).Scan(&teamId).Error; err != nil {
		errors.InternalServerError("something went wrong when searching for users team")
	}

	if err := db.Table("users").Where("id = ?", uid).
		Update("coin", gorm.Expr("coin - ?", cup.EntryPricePerTeam)).
		Error; err != nil {
		return errors.InternalServerError("something went wrong when depositing money from user")
	}

	//TODO : Finish cup
	if err := db.Create(&cup).Error; err != nil {
		return errors.InternalServerError("something went wrong when creating cup")
	}

	if err := db.Table("teams").Where("id = ?", teamId).
		Updates(map[string]interface{}{"cup_id": cup.ID}).Error; err != nil {
		return errors.InternalServerError("something went wrong when setting cup id")
	}

	return nil
}

func (cup *Cup) JoinCup(uid, cupId int64) *errors.RestErr {
	var teamId uint
	var team Team
	var teamIds []int64
	var matchCount = 0

	cupErr := db.Table("cups").Where("id = ?", cupId).First(cup).Error
	if checkErr.Is(cupErr, gorm.ErrRecordNotFound) {
		return errors.NotFoundError("cup not found")
	}

	userErr := db.Table("users").Select("team_id").Where("id = ?", uid).Scan(&teamId).Error
	if checkErr.Is(userErr, gorm.ErrRecordNotFound) {
		return errors.NotFoundError("user not found")
	}

	teamErr := db.Table("teams").Where("id = ?", teamId).Scan(&team).Error
	if checkErr.Is(teamErr, gorm.ErrRecordNotFound) {
		return errors.NotFoundError("team not found")
	}

	if team.CupId != 0 {
		return errors.BadRequestError("team is already in a cup")
	}

	if err := db.Table("cups").Where("id = ?", cupId).
		Updates(map[string]interface{}{"invites": gorm.Expr("invites + ?", 1)}).Error; err != nil {
		return errors.InternalServerError("something went wrong 1")
	}

	if err := db.Table("users").Where("id = ?", uid).
		Updates(map[string]interface{}{"coin": gorm.Expr("coin - ?", cup.EntryPricePerTeam)}).Error; err != nil {
		return errors.InternalServerError("something went wrong 2")
	}

	if err := db.Table("teams").Where("id = ?", teamId).
		Updates(map[string]interface{}{"cup_id": cupId}).Error; err != nil {
		return errors.InternalServerError("something went wrong 3")
	}

	// TODO : IF MAX_CAPACITY == INVITES ==> CREATE MATCHES.
	if cup.MaxCapacity-cup.Invites == 1 {
		if err := db.Select("id").Table("teams").Where("cup_id = ?", cup.ID).Scan(&teamIds).Error; err != nil {
			return errors.BadRequestError("something went wrong 1")
		}
		fmt.Println(teamIds)
		matchTime := time.Now().AddDate(0, 0, 1).String()

		for i, id := range teamIds {
			i += 1
			match := Match{
				Model:        gorm.Model{},
				CupId:        cup.ID,
				TeamOneID:    uint(id),
				TeamTwoID:    uint(teamIds[i+1]),
				TeamOneGoals: 0,
				TeamTwoGoals: 0,
				WinnerID:     0,
				MatchTime:    matchTime,
				IsDone:       0,
				MatchType:    0,
				IsAccepted:   0,
				Priority:     i,
				MatchTimeDb:  time.Now().AddDate(0, 0, 1),
			}
			fmt.Println("possible matches", id, teamIds[i+1])
			if err := db.Create(&match).Error; err != nil {
				return errors.InternalServerError(err.Error())
			}
			matchCount += 1
			if matchCount == len(teamIds)/2 {
				return nil
			}
		}
		return errors.BadRequestError("cup is fully loaded")
	}

	return nil
}

func (cup *Cup) GetAllCups() ([]Cup, *errors.RestErr) {
	var cups []Cup
	if err := db.Table("cups").Where("name NOT LIKE '%hazfi%'").Find(&cups).Limit(50).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return cups, nil
}

func (cup *Cup) GetOneCup(cupId int64) (*Cup, *errors.RestErr) {
	err := db.Table("cups").Where("id = ?", cupId).First(&cup).Error

	if checkErr.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.NotFoundError("record not found")
	}
	if err != nil {
		return nil, errors.InternalServerError("something went wrong")

	}

	return cup, nil
}

func (cup *Cup) GetCupMatches(cupId int64) ([]map[string]interface{}, *errors.RestErr) {
	var matches []map[string]interface{}
	err := db.Select("team_one_id, team_two_id, team_one_goals, team_two_goals, winner_id, is_done, match_time_db, priority").
		Table("matches").Where("cup_id = ? OR cup_h_id = ?", cupId, cupId).Find(&matches).Error

	if checkErr.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.NotFoundError("record not found")
	}
	if err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return matches, nil
}

func (cup *Cup) UserLeagueAndCupData(uid int64) (map[string]uint, *errors.RestErr) {
	var teamId int64
	var team Team
	values := make(map[string]uint)
	err := db.Select("team_id").Table("users").Where("id = ?", uid).Find(&teamId).Error
	db.Select("league_id, cup_id, cup_h_id").Table("teams").Where("id = ?", teamId).Find(&team)

	values["team id"] = uint(teamId)
	values["league id"] = team.LeagueId
	values["cup h id"] = team.CupHId
	values["cup id"] = team.CupId
	if checkErr.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.NotFoundError("record not found")
	}
	if err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return values, nil
}
