package domain

import (
	checkErr "errors"
	"fmt"
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"github.com/kasrashrz/Affogato/utils/errors"
	"gorm.io/gorm"
	"time"
)

func (playerPower *PlayerPower) PlayerPractice(id int64, paramsKey string) *errors.RestErr {
	query := fmt.Sprintf("%s", paramsKey)
	exprQuery := fmt.Sprintf("%s + ?", paramsKey)
	//var powerToUpdate int
	var trainer dto_only.Trainer

	// check if remaining is -1
	fmt.Println(id)
	if err := db.Select("*").Table("players").
		Joins("join player_powers pp on players.id = pp.player_id").
		Joins("join teams t on players.team_id = t.id").
		Joins("join trainers t2 on t.trainer_id = t2.id").
		Where("players.id = ?", id).Scan(&trainer).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("team does not have a manager")
		}
		return errors.InternalServerError("something went wrong")
	}
	db.Select("*").Table("trainers").
		Joins("join trainer_levels tl on trainers.trainer_level_id = tl.id").
		Where("trainers.id = ?", trainer.ID).Scan(&trainer.TrainerLevel)

	if trainer.RemainingDailyPracticeDuration <= 0 ||
		trainer.RemainingExtraPracticeDuration <= 0 ||
		trainer.RemainingPlayerPracticeDuration <= 0 {
		return errors.BadRequestError("trainer couldn't train today. Limit over")
	}

	//if err := db.Select("skill").Table("players").Where("id = ?", id).
	//	Scan(&powerToUpdate).Error; err != nil {
	//	return errors.InternalServerError("something went wrong with player skill")
	//}

	//if err := db.Select("tl.power").Table("player_powers").
	//	Joins("join players p on p.id = player_powers.player_id").
	//	Joins("join teams t on t.id = p.team_id").
	//	Joins("join trainers t2 on t.trainer_id = t2.id").
	//	Joins("join trainer_levels tl on t2.trainer_level_id = tl.id").
	//	Where("player_id = ?", id).Scan(&powerToUpdate).Error; err != nil {
	//	if checkErr.Is(err, gorm.ErrRecordNotFound) {
	//		return errors.NotFoundError("team does not have a manager")
	//	}
	//	return errors.InternalServerError("something went wrong")
	//}

	if trainer.TrainerLevel.Power == 0 {
		return errors.NotFoundError("team does not have a manager")
	}

	player := Player{}
	findErr := db.Table("players").Where("id = ?", id).First(&player).Error
	if checkErr.Is(findErr, gorm.ErrRecordNotFound) {
		return errors.NotFoundError("player does not exist")
	}

	playerPWR := PlayerPower{}
	playerPowerErr := db.Table("player_powers").Where("player_id = ?", player.ID).First(&playerPWR).Error
	if checkErr.Is(playerPowerErr, gorm.ErrRecordNotFound) {
		return errors.NotFoundError("player power does not exist")
	}

	if playerPWR.Power < 25 || playerPWR.Control < 25 || playerPWR.Dribble < 25 ||
		playerPWR.Endurance < 25 || playerPWR.Goal < 25 || playerPWR.Head < 25 ||
		playerPWR.Pass < 25 || playerPWR.Shoot < 25 {
		return errors.BadRequestError("player powers are less than 25")
	}

	fmt.Println(player.ID)
	if err := db.Table("player_powers").Where("id = ?", playerPWR.ID).
		Update(query, gorm.Expr(exprQuery, trainer.TrainerLevel.Power)).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Exec("update trainers "+
		"join teams t on trainers.id = t.trainer_id "+
		"join players p on t.id = p.team_id "+
		"set remaining_daily_practice_duration= remaining_daily_practice_duration - 1, "+
		"remaining_extra_practice_duration= remaining_extra_practice_duration - 1, "+
		"remaining_player_practice_duration= remaining_player_practice_duration - 1 "+
		"where p.id = ?", id).
		Error; err != nil {
		fmt.Println(err)
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("trainers").Where("id = ?", trainer.ID).
		Updates(map[string]interface{}{"updated_at": time.Now()}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (playerPower *PlayerPower) SetDefaultPowers(id int64) *errors.RestErr {

	if err := db.Table("player_powers").Where("id = ?", id).
		Updates(map[string]interface{}{"control": 25, "power": 25, "pass": 25, "shoot": 25, "dribble": 25,
			"head": 25, "endurance": 25, "strength": 25, "goal": 25}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (playerPower *PlayerPower) CheckTrainers() *errors.RestErr {
	var trainers []dto_only.Trainer
	var power int
	if err := db.Table("trainers").Where("hour(timediff(now(), updated_at)) > 23").
		Scan(&trainers).Error; err != nil {
		fmt.Println(err)
	}

	for _, trainer := range trainers {
		db.Select("tl.power").Table("trainers").
			Joins("join trainer_levels tl on trainers.trainer_level_id = tl.id").
			Where("trainers.id = ?", trainer.ID).
			Scan(&power)
		if err := db.Table("trainers").
			Where("id = ?", trainer.ID).
			Updates(map[string]interface{}{"remaining_daily_practice_duration": power,
				"remaining_player_practice_duration": power,
				"remaining_extra_practice_duration":  power,
				"updated_at":                         time.Now()}).Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}
	}

	power = 0

	return nil
}
