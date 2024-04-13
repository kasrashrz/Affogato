package domain

import (
	checkErr "errors"
	"fmt"
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"github.com/kasrashrz/Affogato/utils/errors"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (team *GetAllTeams) GetAll() ([]GetAllTeams, *errors.RestErr) {

	var teams []GetAllTeams
	if err := db.Table("teams").Limit(100).Scan(&teams).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return teams, nil
}

func (team *Team) TeamsCount() int {
	var teamsCount int
	if err := db.Select("count(id)").Table("teams").Where("league_id IS NULL").Scan(&teamsCount).Error; err != nil {
		return 0
	}
	//logger.Info(fmt.Sprintf("Team Counts: %s", teamsCount))
	return teamsCount
}

func (team *GetAllTeams) GetOne(uid int64) (*Team, *errors.RestErr) {
	var outputTeam Team

	if err := db.Select("team_id").Table("users").
		Where("id = ?", uid).Scan(&team.ID).Error; err != nil {
		if strings.Contains(err.Error(), "converting NULL to uint is unsupported") {
			return nil, errors.BadRequestError("user does not have team")
		}
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Preload("Stadium.Ground").
		Preload("Stadium.ScoreBoard").
		Preload("Stadium.Lights").
		Preload("AssistantCoach").
		Preload("Doctor").
		Preload("TalentFinder").
		Preload("Trainer").
		Preload("FitnessCoach").
		Preload("Tactic").
		Where("id = ?", team.ID).
		First(&outputTeam).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, errors.NotFoundError("team does not exist")
		}
		return nil, errors.InternalServerError("something went wrong")
	}
	return &outputTeam, nil
}

func (inpTeam *Team) Create(uid int64) *errors.RestErr {
	rand.Seed(time.Now().UnixNano())
	var TrainerIds []int
	if err := db.Table("trainers").Select("id").Where("trainer_level_id = ?", 1).
		Scan(&TrainerIds).
		Error; err != nil {

	}

	assistantCoach := dto_only.AssistantCoach{
		Model:                 gorm.Model{},
		AssistantCoachLevel:   levels.AssistantCoachLevel{},
		AssistantCoachLevelId: 1,
		Practices:             3,
	}
	trainer := dto_only.Trainer{
		Model:                           gorm.Model{},
		ContractExp:                     "",
		WeeklySalary:                    0,
		TrainerLevel:                    levels.TrainerLevel{},
		TrainerLevelId:                  1,
		RemainingDailyPracticeDuration:  3,
		RemainingPlayerPracticeDuration: 3,
		RemainingExtraPracticeDuration:  3,
	}
	fitnessCoach := dto_only.FitnessCoach{
		Model:               gorm.Model{},
		WeeklySalary:        0,
		FitnessCoachLevel:   levels.FitnessCoachLevel{},
		FitnessCoachLevelId: 1,
	}
	talentFinder := dto_only.TalentFinder{
		Model:               gorm.Model{},
		WeeklySalary:        0,
		TalentFinderLevel:   levels.TalentFinderLevel{},
		TalentFinderLevelId: 1,
	}
	doctor := dto_only.Doctor{
		Model:         gorm.Model{},
		DoctorLevel:   levels.DoctorLevel{},
		DoctorLevelId: 1,
	}

	db.Create(&assistantCoach)
	db.Create(&fitnessCoach)
	db.Create(&trainer)
	db.Create(&talentFinder)
	db.Create(&doctor)

	team := Team{
		Model:            gorm.Model{},
		Stadium:          dto_only.Stadium{},
		Name:             inpTeam.Name,
		TrainerId:        trainer.ID,
		AssistantCoachId: assistantCoach.ID,
		FitnessCoachId:   fitnessCoach.ID,
		DoctorId:         doctor.ID,
		TalentFinderId:   talentFinder.ID,
		Tactic:           dto_only.Tactic{},
		TacticId:         1,
		TacticFormation:  "175M2-177D4-181A8-182M8-200A3",
	}

	team.LastEarned = time.Now()
	team.LastPaid = time.Now()
	if err := db.Create(&team).Error; err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "Error 1452") {
			if strings.Contains(err.Error(), "league") {
				return errors.BadRequestError("league does not exist")
			}
			if strings.Contains(err.Error(), "tactic") {
				return errors.BadRequestError("tactic does not exist")
			}
		}
	}

	stadium := dto_only.Stadium{
		Model:          gorm.Model{},
		TotalCapacity:  100,
		TicketPrice:    1000,
		StadiumLevelId: 1,
		Pleasure:       0,
		Ground: dto_only.Ground{
			Model:         gorm.Model{},
			GroundLevel:   levels.GroundLevels{},
			GroundLevelID: 1,
			RequiredScore: 1,
			StadiumID:     team.Stadium.ID,
		},
		ScoreBoard: dto_only.ScoreBoard{},
		Lights:     dto_only.Lights{},
		TeamID:     team.ID,
	}
	if err := db.Create(&stadium).Error; err != nil {
		if strings.Contains(err.Error(), "Error 1452") {
			if strings.Contains(err.Error(), "league") {
				return errors.BadRequestError("league does not exist")
			}
			if strings.Contains(err.Error(), "tactic") {
				return errors.BadRequestError("tactic does not exist")
			}
		}
		return errors.InternalServerError("something went wrong")
	}
	scoreBoard := dto_only.ScoreBoard{
		Model:             gorm.Model{},
		ScoreBoardLevelID: 1,
		Price:             0,
		RequiredScore:     0,
		StadiumID:         stadium.ID,
	}
	lights := dto_only.Lights{
		Model:         gorm.Model{},
		LightsLevel:   levels.LightLevels{},
		LightsLevelId: 1,
		StadiumId:     stadium.ID,
		Price:         0,
		RequiredScore: 0,
		IncreaseFans:  0,
	}
	if err := db.Save(&scoreBoard).Error; err != nil {
		if strings.Contains(err.Error(), "score_board_levels") {
			return errors.NotFoundError("score board not found")
		}
		return errors.InternalServerError("something went wrong")
	}
	if err := db.Save(&lights).Error; err != nil {
		if strings.Contains(err.Error(), "light_levels") {
			return errors.NotFoundError("light levels not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	restaurant := dto_only.Restaurant{
		Model:             gorm.Model{},
		RestaurantLevel:   levels.RestaurantLevel{},
		RestaurantLevelId: 1,
		StadiumID:         stadium.ID,
		LastPaid:          time.Now(),
		LastEarned:        time.Now(),
	}
	shopping := dto_only.Shopping{
		Model:           gorm.Model{},
		ShoppingLevel:   levels.ShoppingLevel{},
		ShoppingLevelId: 1,
		StadiumID:       stadium.ID,
		LastPaid:        time.Now(),
		LastEarned:      time.Now(),
	}
	parking := dto_only.Parking{
		Model:          gorm.Model{},
		ParkingLevel:   levels.ParkingLevel{},
		ParkingLevelId: 1,
		StadiumID:      stadium.ID,
		LastPaid:       time.Now(),
		LastEarned:     time.Now(),
	}
	transportation := dto_only.Transportation{
		Model:                 gorm.Model{},
		TransportationLevel:   levels.TransportationLevel{},
		TransportationLevelId: 1,
		StadiumID:             stadium.ID,
		LastPaid:              time.Now(),
		LastEarned:            time.Now(),
	}

	if err := db.Save(&restaurant).Error; err != nil {
		if strings.Contains(err.Error(), "restaurant_levels") {
			return errors.NotFoundError("restaurant levels not found")
		}
		return errors.InternalServerError("something went wrong")
	}
	if err := db.Save(&shopping).Error; err != nil {
		if strings.Contains(err.Error(), "shopping_levels") {
			return errors.NotFoundError("shopping levels not found")
		}
		return errors.InternalServerError("something went wrong")
	}
	if err := db.Save(&parking).Error; err != nil {
		if strings.Contains(err.Error(), "parking_levels") {
			return errors.NotFoundError("parking levels not found")
		}
		return errors.InternalServerError("something went wrong")
	}
	if err := db.Save(&transportation).Error; err != nil {
		if strings.Contains(err.Error(), "transportation_levels") {
			return errors.NotFoundError("transportation levels not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("users").Where("id = ?", uid).
		Updates(map[string]interface{}{"team_id": team.ID}).Error; err != nil {
		return errors.InternalServerError("something went wrong when assigning team_id to user")
	}

	return nil
}

func (inpTeam *Team) CreateRandom(quantity int) *errors.RestErr {
	rand.Seed(time.Now().UnixNano())
	//var TrainerIds []int

	//if err := db.Table("trainers").Select("id").Where("trainer_level_id = ?", 1).
	//	Scan(&TrainerIds).
	//	Error; err != nil {
	//
	//}

	for i := 0; i < quantity; i++ {
		assistantCoach := dto_only.AssistantCoach{
			Model:                 gorm.Model{},
			AssistantCoachLevel:   levels.AssistantCoachLevel{},
			AssistantCoachLevelId: 1,
			Practices:             3,
		}
		trainer := dto_only.Trainer{
			Model:                           gorm.Model{},
			ContractExp:                     "",
			WeeklySalary:                    0,
			TrainerLevel:                    levels.TrainerLevel{},
			TrainerLevelId:                  1,
			RemainingDailyPracticeDuration:  3,
			RemainingPlayerPracticeDuration: 3,
			RemainingExtraPracticeDuration:  3,
		}
		fitnessCoach := dto_only.FitnessCoach{
			Model:               gorm.Model{},
			WeeklySalary:        0,
			FitnessCoachLevel:   levels.FitnessCoachLevel{},
			FitnessCoachLevelId: 1,
		}
		talentFinder := dto_only.TalentFinder{
			Model:               gorm.Model{},
			WeeklySalary:        0,
			TalentFinderLevel:   levels.TalentFinderLevel{},
			TalentFinderLevelId: 1,
		}
		doctor := dto_only.Doctor{
			Model:         gorm.Model{},
			DoctorLevel:   levels.DoctorLevel{},
			DoctorLevelId: 1,
		}
		db.Create(&assistantCoach)
		db.Create(&fitnessCoach)
		db.Create(&trainer)
		db.Create(&talentFinder)
		db.Create(&doctor)

		team := Team{
			Model:            gorm.Model{},
			AssistantCoach:   assistantCoach,
			AssistantCoachId: assistantCoach.ID,
			FitnessCoach:     fitnessCoach,
			FitnessCoachId:   fitnessCoach.ID,
			Doctor:           doctor,
			DoctorId:         doctor.ID,
			TalentFinder:     talentFinder,
			TalentFinderId:   talentFinder.ID,
			Trainer:          trainer,
			TrainerId:        trainer.ID,
			Name:             "random team " + strconv.Itoa(rand.Intn(10000)),
			TacticFormation:  "175M2-177D4-181A8-182M8-200A3",
		}

		team.AssistantCoachId = assistantCoach.ID
		team.DoctorId = team.ID
		team.TalentFinderId = talentFinder.ID
		team.TrainerId = trainer.ID
		team.FitnessCoachId = fitnessCoach.ID

		if err := db.Create(&team).Error; err != nil {
			fmt.Println(err)
			if strings.Contains(err.Error(), "Error 1452") {
				if strings.Contains(err.Error(), "league") {
					return errors.BadRequestError("league does not exist")
				}
				if strings.Contains(err.Error(), "tactic") {
					return errors.BadRequestError("tactic does not exist")
				}
			}
		}

	}

	return nil
}

func (team *Team) FindPlayers(id int64) ([]Player, *errors.RestErr) {
	var players []Player

	if err := db.
		Table("players").
		Preload("Power").
		Preload("Post").
		Preload("Status").
		Joins("join teams t on t.id = players.team_id").
		Joins("join users u on t.id = u.team_id").
		Where("u.id = ?", id).
		Find(&players).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return players, nil
}

func (team *Team) TeamPowerAvg(id int64) (avg int64, err *errors.RestErr) {
	var queryAvg int64
	averageQuery := "ROUND(AVG(shoot + pass + strength + endurance + dribble + control + head + pp.goal + pp.power + tackle))"
	db.Select(averageQuery).Table("players").
		Joins("inner join player_powers pp on players.id = pp.player_id").
		Where("players.team_id = ?", id).Scan(&queryAvg)
	avg = queryAvg
	return avg, nil
}

func (team *Team) TeamEnergyAvg(id int64) (avg int64, err *errors.RestErr) {

	averageQuery := "ROUND(cast(AVG(energy) as float ))"
	if err := db.Select(averageQuery).Table("players").
		Joins("inner join player_powers pp on players.id = pp.player_id").
		Where("players.team_id = ?", id).
		Scan(&avg).Error; err != nil {

		if avg == 0 {
			return 0, errors.NotFoundError("team energy average is 0")
		}
		return 0, errors.InternalServerError("something went wrong")
	}
	return avg, nil
}

func (team *Team) Delete(id int64) *errors.RestErr {

	if err := db.Exec("delete teams, s, g from teams"+
		" join stadia s on teams.id = s.team_id"+
		" join grounds g on s.id = g.stadium_id"+
		" where teams.id = ?", id).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (team *Team) ChangeScore(id, score int64) *errors.RestErr {

	err := db.Table("teams").Where("id = ?", id).First(&team).Error

	if checkErr.Is(err, gorm.ErrRecordNotFound) {
		return errors.NotFoundError("team does not exist")
	}
	if score < 0 {
		if team.Score+score < 0 {
			return errors.BadRequestError("not enough score to decrease")
		}
	}
	if err := db.Table("teams").Where("id = ?", id).
		Update("score", gorm.Expr("score + ?", score)).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (team *Team) ChangeName(uid int64, name string) *errors.RestErr {
	var teamId int64

	if err := db.Select("t.id").
		Table("users").
		Joins("join teams t on t.id = users.team_id").
		Where("users.id = ?", uid).
		Scan(&teamId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").Where("id = ?", teamId).First(&team).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("Team does not exist")
		}
	}
	if team.ChangedName != false {
		return errors.BadRequestError("team name has already changed")
	}
	if err := db.Table("teams").Where("id = ?", teamId).
		Updates(map[string]interface{}{"name": name, "changed_name": true}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (team *Team) RandomGenerate(quantity int64) *errors.RestErr {
	return nil
}

func (team *Team) TeamTactic(uid, opponentId int64) (map[string]string, *errors.RestErr) {

	var assistantCoachId uint
	var tactic dto_only.Tactic
	var userTeamId, opponentTeamId int64
	var tacticFormation string
	tacticData := make(map[string]string)
	if err := db.Select("t.id").
		Table("users").
		Joins("join teams t on t.id = users.team_id").
		Where("users.id = ?", uid).
		Scan(&userTeamId).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("t.id").
		Table("users").
		Joins("join teams t on t.id = users.team_id").
		Where("users.id = ?", opponentId).
		Scan(&opponentTeamId).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("assistant_coach_id").Table("teams").
		Where("id = ?", userTeamId).Scan(&assistantCoachId).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if assistantCoachId == 0 {
		return nil, errors.BadRequestError("your team does not have coach")
	}

	if err := db.Table("tactics").Select("tactics.name, tactic_format, strategyq").
		Joins("join teams t on tactics.id = t.tactic_id").
		Where("t.id = ?", opponentTeamId).Scan(&tactic).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	if err := db.Table("teams").Select("tactic_formation").
		Where("id = ?", opponentTeamId).Scan(&tacticFormation).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	tacticData["tactic-name"] = tactic.Name
	tacticData["tactic-strategy"] = tactic.Strategy
	tacticData["tactic-format"] = tactic.TacticFormat
	tacticData["tactic-formation"] = tacticFormation

	return tacticData, nil
}

func (team *Team) Workers(userId int64) (
	*dto_only.AssistantCoach, *dto_only.Trainer,
	*dto_only.FitnessCoach, *dto_only.Doctor,
	*dto_only.TalentFinder, *errors.RestErr) {

	var trainer dto_only.Trainer
	var assistantCoach dto_only.AssistantCoach
	var fitnessCoach dto_only.FitnessCoach
	var doctor dto_only.Doctor
	var talentFinder dto_only.TalentFinder

	if err := db.Preload("TrainerLevel").Table("users").Select("*").
		Joins("join teams t on users.team_id = t.id").
		Joins("join trainers t2 on t.trainer_id = t2.id").
		Where("users.id = ?", userId).
		Find(&trainer).Error; err != nil {
		return nil, nil, nil, nil, nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("AssistantCoachLevel").Table("assistant_coaches").
		Joins("join teams t on assistant_coaches.id = t.assistant_coach_id").
		Joins("join users u on t.id = u.team_id").
		Where("u.id = ?", userId).
		Find(&assistantCoach).Error; err != nil {
		return nil, nil, nil, nil, nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("FitnessCoachLevel").Table("fitness_coaches").
		Joins("join teams t on fitness_coaches.id = t.fitness_coach_id").
		Joins("join users u on t.id = u.team_id").
		Where("u.id = ?", userId).
		Find(&fitnessCoach).Error; err != nil {
		return nil, nil, nil, nil, nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("DoctorLevel").Table("users").Select("*").
		Joins("join teams t on t.id = users.team_id").
		Joins("join doctors d on t.doctor_id = d.id").
		Where("users.id = ?", userId).
		Find(&doctor).Error; err != nil {
		return nil, nil, nil, nil, nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("TalentFinderLevel").Table("talent_finders").
		Joins("join teams t on talent_finders.id = t.talent_finder_id").
		Joins("join users u on t.id = u.team_id").
		Where("u.id = ?", userId).
		Find(&talentFinder).Error; err != nil {
		return nil, nil, nil, nil, nil, errors.InternalServerError("something went wrong")
	}

	return &assistantCoach, &trainer, &fitnessCoach, &doctor, &talentFinder, nil
}

func (team *Team) BuyAssistantCoach(userId int64, newAssistantCoachId int64) *errors.RestErr {
	var user User
	var currentAssistantCoach dto_only.AssistantCoach
	var teamData Team

	if err := db.Table("users").
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("user not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", user.TeamId).
		First(&teamData).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("team not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	db.Table("assistant_coaches").Where("id = ?", teamData.AssistantCoachId).First(&currentAssistantCoach)

	newAssistantCoachId = int64(currentAssistantCoach.AssistantCoachLevelId) + 1

	if currentAssistantCoach.AssistantCoachLevelId == 6 {
		return errors.BadRequestError("you already have highest assistant coach level")
	}

	paymentData := PaymentHistory{
		Model:            gorm.Model{},
		Team:             Team{},
		TeamId:           int64(user.TeamId),
		PaymentDetail:    PaymentDetail{},
		PaymentDetailId:  1,
		AssistantCoach:   dto_only.AssistantCoach{},
		AssistantCoachId: newAssistantCoachId,
	}

	if err := db.Create(&paymentData).Error; err != nil {
		return errors.InternalServerError("something went wrong while saving payment")
	}

	if err := db.Table("assistant_coaches").
		Where("id = ?", currentAssistantCoach.ID).
		Update("assistant_coach_level_id", gorm.Expr("assistant_coach_level_id + ?", 1)).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	event := "assistant coach level " + strconv.FormatInt(newAssistantCoachId, 10) + " joined your team"
	if err := team.SubmitTransferData(userId, 0, event); err != nil {
		return errors.InternalServerError("something went wrong setting transfer data")
	}

	return nil

}

func (team *Team) BuyDoctor(userId int64, newDoctorId int64) *errors.RestErr {
	var user User
	var currentDoctor dto_only.Doctor
	var teamData Team

	if err := db.Table("users").
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("user not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", user.TeamId).
		First(&teamData).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("team not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Joins("join doctors d on teams.doctor_id = d.id").
		Where("teams.id = ?", teamData.ID).
		Scan(&currentDoctor).
		Error; err != nil {

	}

	newDoctorId = int64(currentDoctor.DoctorLevelId) + 1

	if currentDoctor.DoctorLevelId == 6 {
		return errors.BadRequestError("you already have highest doctor level")
	}

	if currentDoctor.ID == uint(newDoctorId) {
		return errors.BadRequestError("already hired the doctor")
	}

	paymentData := PaymentHistory{
		Model:           gorm.Model{},
		Team:            Team{},
		TeamId:          int64(user.TeamId),
		PaymentDetail:   PaymentDetail{},
		PaymentDetailId: 2,
		Doctor:          dto_only.Doctor{},
		DoctorId:        newDoctorId,
	}

	if err := db.Create(&paymentData).Error; err != nil {
		return errors.InternalServerError("something went wrong while saving payment")
	}

	if err := db.Table("doctors").
		Where("id = ?", currentDoctor.ID).
		Update("doctor_level_id", gorm.Expr("doctor_level_id + ?", 1)).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	event := "doctor level " + strconv.FormatInt(newDoctorId, 10) + " joined your team"
	if err := team.SubmitTransferData(userId, 0, event); err != nil {
		return errors.InternalServerError("something went wrong setting transfer data")
	}

	return nil
}

func (team *Team) BuyFitnessCoach(userId int64, newFitnessCoachId int64) *errors.RestErr {
	var user User
	var currentFitnessCoach dto_only.FitnessCoach
	var teamData Team

	if err := db.Table("users").
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("user not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", user.TeamId).
		First(&teamData).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("team not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	db.Table("fitness_coaches").
		Where("id = ?", teamData.FitnessCoachId).
		First(&currentFitnessCoach)

	if currentFitnessCoach.ID == uint(newFitnessCoachId) {
		return errors.BadRequestError("already hired the fitness coach")
	}

	newFitnessCoachId = int64(currentFitnessCoach.FitnessCoachLevelId) + 1

	if currentFitnessCoach.FitnessCoachLevelId == 6 {
		return errors.BadRequestError("you already have highest fitness coach")
	}

	paymentData := PaymentHistory{
		Model:           gorm.Model{},
		Team:            Team{},
		TeamId:          int64(user.TeamId),
		PaymentDetail:   PaymentDetail{},
		PaymentDetailId: 3,
		FitnessCoach:    dto_only.FitnessCoach{},
		FitnessCoachId:  newFitnessCoachId,
	}

	if err := db.Create(&paymentData).Error; err != nil {
		return errors.InternalServerError("something went wrong while saving payment")
	}

	if err := db.Table("fitness_coaches").
		Where("id = ?", currentFitnessCoach.ID).
		Update("fitness_coach_level_id", gorm.Expr("fitness_coach_level_id + ?", 1)).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	event := "fitness coach level " + strconv.FormatInt(newFitnessCoachId, 10) + " joined your team"
	if err := team.SubmitTransferData(userId, 0, event); err != nil {
		return errors.InternalServerError("something went wrong setting transfer data")
	}

	return nil
}

func (team *Team) BuyTalentFinder(userId int64, newTalentFinderId int64) *errors.RestErr {
	var user User
	var currentTalentFinder dto_only.TalentFinder
	var teamData Team

	if err := db.Table("users").
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("user not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", user.TeamId).
		First(&teamData).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("team not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("talent_finders").
		Where("id = ?", teamData.TalentFinderId).
		Scan(&currentTalentFinder).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("talent finder not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	newTalentFinderId = int64(currentTalentFinder.TalentFinderLevelId) + 1

	if currentTalentFinder.TalentFinderLevelId == 6 {
		return errors.BadRequestError("you already have highest talent finder")
	}

	if err := db.Table("talent_finders").
		Where("id = ?", currentTalentFinder.ID).
		Update("talent_finder_level_id", gorm.Expr("talent_finder_level_id + ?", 1)).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	paymentData := PaymentHistory{
		Model:           gorm.Model{},
		Team:            Team{},
		TeamId:          int64(user.TeamId),
		PaymentDetail:   PaymentDetail{},
		PaymentDetailId: 4,
		TalentFinder:    dto_only.TalentFinder{},
		TalentFinderId:  newTalentFinderId,
	}

	if err := db.Create(&paymentData).Error; err != nil {
		return errors.InternalServerError("something went wrong while saving payment")
	}

	event := "talent finder level " + strconv.FormatInt(newTalentFinderId, 10) + " joined your team"
	if err := team.SubmitTransferData(userId, 0, event); err != nil {
		return errors.InternalServerError("something went wrong setting transfer data")
	}

	return nil
}

func (team *Team) BuyTrainer(userId int64, newTrainerId int64) *errors.RestErr {
	var user User
	var currentTrainer dto_only.Trainer
	var teamData Team

	if err := db.Table("users").
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("user not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", user.TeamId).
		First(&teamData).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("team not found")
		}
		return errors.InternalServerError("something went wrong")
	}
	db.Table("trainers").Where("id = ?", teamData.TrainerId).
		First(&currentTrainer)

	newTrainerId = int64(currentTrainer.TrainerLevelId) + 1
	if currentTrainer.TrainerLevelId == 6 {
		return errors.BadRequestError("you already have highest trainer level")
	}

	if currentTrainer.ID == uint(newTrainerId) {
		return errors.BadRequestError("already hired the trainer")
	}

	paymentData := PaymentHistory{
		Model:           gorm.Model{},
		Team:            Team{},
		TeamId:          int64(user.TeamId),
		PaymentDetail:   PaymentDetail{},
		PaymentDetailId: 5,
		Trainer:         dto_only.Trainer{},
		TrainerId:       newTrainerId,
	}

	if err := db.Create(&paymentData).Error; err != nil {
		return errors.InternalServerError("something went wrong while saving payment")
	}

	if err := db.Table("trainers").
		Where("id = ?", currentTrainer.ID).
		Update("trainer_level_id", gorm.Expr("trainer_level_id + ?", 1)).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	event := "trainer level " + strconv.FormatInt(newTrainerId, 10) + " joined your team"
	if err := team.SubmitTransferData(userId, 0, event); err != nil {
		return errors.InternalServerError("something went wrong setting transfer data")
	}

	return nil
}

func (team *Team) GetPurchaseHistory(userId int64) (*[]PaymentHistory, *errors.RestErr) {
	var purchases []PaymentHistory
	var finalPurchases []PaymentHistory
	var purchasDetail PaymentDetail
	var teamId uint
	db.Select("team_id").Table("users").Where("id = ?", userId).Scan(&teamId)

	if err := db.Preload("PaymentDetail").
		Select("payment_histories.id,"+
			"payment_histories.payment_detail_id,"+
			"payment_histories.created_at,"+
			"payment_histories.team_id,"+
			"payment_histories.trainer_id,"+
			"payment_histories.assistant_coach_id,"+
			"payment_histories.doctor_id,"+
			"payment_histories.fitness_coach_id,"+
			"payment_histories.deleted_at,"+
			"payment_histories.talent_finder_id").
		Table("payment_histories").
		Where("team_id = ?", teamId).
		Scan(&purchases).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	if purchases == nil {
		return nil, errors.NotFoundError("user or purchase not found")
	}

	for _, purchase := range purchases {
		db.Table("payment_details").Where("id = ?", purchase.PaymentDetailId).Scan(&purchasDetail)
		purchase.PaymentDetail = purchasDetail
		finalPurchases = append(finalPurchases, purchase)
		purchasDetail = PaymentDetail{}
	}

	return &finalPurchases, nil
}

func (team *Team) GetAllDoctors() (*[]dto_only.Doctor, *errors.RestErr) {

	var docs []dto_only.Doctor

	if err := db.Table("doctors").
		Preload("DoctorLevel").
		Find(&docs).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return &docs, nil
}

func (team *Team) GetAllTrainers() (*[]dto_only.Trainer, *errors.RestErr) {

	var trainers []dto_only.Trainer

	if err := db.Select("*").
		Preload("TrainerLevel").
		Table("trainers").
		Find(&trainers).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return &trainers, nil
}

func (team *Team) GetAllAssistantCoaches() (*[]dto_only.AssistantCoach, *errors.RestErr) {

	var assistantCoaches []dto_only.AssistantCoach

	if err := db.Select("*").
		Preload("AssistantCoachLevel").
		Table("assistant_coaches").
		Find(&assistantCoaches).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return &assistantCoaches, nil
}

func (team *Team) GetAllFitnessCoaches() (*[]dto_only.FitnessCoach, *errors.RestErr) {

	var fitnessCoaches []dto_only.FitnessCoach

	if err := db.Select("*").
		Preload("FitnessCoachLevel").
		Table("fitness_coaches").
		Find(&fitnessCoaches).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return &fitnessCoaches, nil
}

func (team *Team) GetAllTalentFinders() (*[]dto_only.TalentFinder, *errors.RestErr) {

	var talentFinders []dto_only.TalentFinder

	if err := db.Select("*").
		Preload("TalentFinderLevel").
		Table("talent_finders").
		Find(&talentFinders).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return &talentFinders, nil
}

func (team *Team) ChangeTactic(uid, tacticId int64) *errors.RestErr {

	var assistantCoach dto_only.AssistantCoach
	var tactic dto_only.Tactic
	var teamId int64

	if err := db.Table("tactics").Where("id = ?", tacticId).First(&tactic).Error; err != nil {
		if checkErr.Is(gorm.ErrRecordNotFound, err) {
			return errors.NotFoundError("tactic not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("assistant_coaches").
		Joins("join teams t on assistant_coaches.id = t.assistant_coach_id").
		Joins("join users u on t.id = u.team_id").
		Where("u.id = ?", uid).
		Scan(&assistantCoach).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if assistantCoach.ID == 0 {
		return errors.BadRequestError("team does not have an assistant coach")
	}

	if err := db.Select("teams.id").Table("teams").
		Joins("join users u on teams.id = u.team_id").
		Where("u.id = ?", uid).
		Scan(&teamId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if teamId == 0 {
		return errors.BadRequestError("user does not have a team")
	}

	if err := db.Table("teams").
		Where("id = ?", teamId).
		Updates(map[string]interface{}{"tactic_id": tacticId}).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (team *Team) GetStadiumData(uid int64) (map[string]interface{}, *errors.RestErr) {
	var stadium dto_only.Stadium
	var lights dto_only.Lights
	var scoreBoard dto_only.ScoreBoard
	var ground dto_only.Ground
	var parking dto_only.Parking
	var restaurant dto_only.Restaurant
	var shopping dto_only.Shopping
	var transportation dto_only.Transportation
	stadiumData := make(map[string]interface{})

	if err := db.Table("stadia").Where("u.id = ?", uid).
		Joins("join teams t on stadia.team_id = t.id").
		Joins("join users u on t.id = u.team_id").
		Find(&stadium).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("LightsLevel").Table("lights").Where("stadium_id = ?", stadium.ID).
		Find(&lights).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("ScoreBoardLevel").Table("score_boards").Where("stadium_id = ?", stadium.ID).
		Find(&scoreBoard).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("GroundLevel").Table("grounds").Where("stadium_id = ?", stadium.ID).
		Find(&ground).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("ParkingLevel").Table("parkings").Where("stadium_id = ?", stadium.ID).
		Find(&parking).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("RestaurantLevel").Table("restaurants").Where("stadium_id = ?", stadium.ID).
		Find(&restaurant).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("ShoppingLevel").Table("shoppings").Where("stadium_id = ?", stadium.ID).
		Find(&shopping).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Preload("TransportationLevel").Table("transportations").Where("stadium_id = ?", stadium.ID).
		Find(&transportation).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	stadiumData["stadium"] = stadium
	stadiumData["lights"] = lights
	stadiumData["scoreboard"] = scoreBoard
	stadiumData["ground"] = ground
	stadiumData["parking"] = parking
	stadiumData["restaurant"] = restaurant
	stadiumData["shopping"] = shopping
	stadiumData["transportation"] = transportation

	return stadiumData, nil
}

func (team *Team) GetStadiumLevels() (map[string]interface{}, *errors.RestErr) {

	var lightsLevels []levels.LightLevels
	var scoreBoardsLevels []levels.ScoreBoardLevels
	var groundLevels []levels.GroundLevels
	var parkingLevels []levels.GroundLevels
	var restaurantLevels []levels.GroundLevels
	var shoppingLevels []levels.GroundLevels
	var transportationLevels []levels.GroundLevels
	var stadiumLevelsSlice []levels.StadiumLevel
	stadiumLevels := make(map[string]interface{})

	db.Table("light_levels").Scan(&lightsLevels)
	db.Table("ground_levels").Scan(&groundLevels)
	db.Table("score_board_levels").Scan(&scoreBoardsLevels)
	db.Table("parking_levels").Scan(&parkingLevels)
	db.Table("restaurant_levels").Scan(&restaurantLevels)
	db.Table("shopping_levels").Scan(&shoppingLevels)
	db.Table("transportation_levels").Scan(&transportationLevels)
	db.Table("stadium_levels").Scan(&stadiumLevelsSlice)

	stadiumLevels["light levels"] = lightsLevels
	stadiumLevels["ground levels"] = groundLevels
	stadiumLevels["scoreboard levels"] = scoreBoardsLevels
	stadiumLevels["parking levels"] = parkingLevels
	stadiumLevels["restaurant levels"] = restaurantLevels
	stadiumLevels["shopping levels"] = shoppingLevels
	stadiumLevels["transportation levels"] = transportationLevels
	stadiumLevels["stadium levels"] = stadiumLevelsSlice

	return stadiumLevels, nil
}

func (team *Team) Payments() *errors.RestErr {
	var earnedTeamIds []uint
	var paidTeamIds []uint
	type salaryPayList struct {
		ID     uint  `json:"id"`
		TeamId uint  `json:"team_id"`
		Sum    int64 `json:"sum"`
	}
	var payList []salaryPayList
	var playerPayList []salaryPayList
	if err := db.Select("id").Table("teams").
		Where("last_earned IS NULL OR hour(timediff(now(), last_earned)) > 23").
		Scan(&earnedTeamIds).Error; err != nil {
		return errors.InternalServerError("something went wrong payment")
	}

	if err := db.Select("id").Table("teams").
		Where("last_paid IS NULL OR hour(timediff(now(), last_paid)) > 167").
		Scan(&paidTeamIds).Error; err != nil {
		return errors.InternalServerError("something went wrong payment")
	}

	for _, teamId := range earnedTeamIds {
		incomePrice := 0
		db.Select("rl.daily_income_price + sl.daily_income_price + pl.daily_income_price").Table("stadia").
			Joins("join restaurants r on stadia.id = r.stadium_id").
			Joins("join restaurant_levels rl on r.restaurant_level_id = rl.id").
			Joins("join shoppings s on stadia.id = s.stadium_id").
			Joins("join shopping_levels sl on s.shopping_level_id = sl.id").
			Joins("join parkings p on stadia.id = p.stadium_id").
			Joins("join parking_levels pl on p.parking_level_id = pl.id").Where("team_id = ?", teamId).Scan(&incomePrice)
		if err := db.Table("users").
			Where("team_id = ?", teamId).
			Update("coin", gorm.Expr("coin + ?", incomePrice)).
			Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}
		db.Table("teams").Where("id = ?", teamId).Updates(map[string]interface{}{"last_earned": time.Now()})
	}

	db.Select("u.id, u.team_id, (acl.salary + tl.salary + dl.salary + tfl.salary + fcl.salary) as sum").Table("teams").
		Joins("join users u on teams.id = u.team_id").
		Joins("join assistant_coaches ac on ac.id = teams.assistant_coach_id").
		Joins("join assistant_coach_levels acl on acl.id = ac.assistant_coach_level_id").
		Joins("join trainers t2 on teams.trainer_id = t2.id").
		Joins("join trainer_levels tl on t2.trainer_level_id = tl.id").
		Joins("join doctors d on teams.doctor_id = d.id").
		Joins("join doctor_levels dl on d.doctor_level_id = dl.id").
		Joins("join talent_finders tf on teams.talent_finder_id = tf.id").
		Joins("join talent_finder_levels tfl on tf.talent_finder_level_id = tfl.id").
		Joins("join fitness_coaches fc on fc.id = teams.fitness_coach_id").
		Joins("join fitness_coach_levels fcl on fcl.id = fc.fitness_coach_level_id").
		Where("hour(timediff(now(), teams.last_paid)) > 167").Scan(&payList)

	db.Select("t.id as team_id, SUM(players.price) AS sum").Table("players").
		Joins("JOIN teams t ON players.team_id = t.id").
		Where("hour(timediff(now(), t.last_paid)) > 167 GROUP BY team_id;").Scan(&playerPayList)

	for _, payment := range payList {
		if err := db.Table("users").
			Where("id = ?", payment.ID).
			Update("coin", gorm.Expr("coin - ?", payment.Sum)).
			Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}
		db.Table("teams").Where("id = ?", payment.TeamId).Updates(map[string]interface{}{"last_paid": time.Now()})
		paymentHistory := PaymentHistory{
			Model:           gorm.Model{},
			Team:            Team{},
			TeamId:          int64(payment.TeamId),
			PaymentDetail:   PaymentDetail{},
			Price:           payment.Sum,
			PaymentDetailId: 7,
		}
		db.Save(&paymentHistory)
	}

	for _, payment := range playerPayList {
		if err := db.Table("users").
			Where("id = ?", payment.ID).
			Update("coin", gorm.Expr("coin - ?", payment.Sum)).
			Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}
		db.Table("teams").Where("id = ?", payment.TeamId).Updates(map[string]interface{}{"last_paid": time.Now()})
		paymentHistory := PaymentHistory{
			Model:           gorm.Model{},
			Team:            Team{},
			TeamId:          int64(payment.TeamId),
			PaymentDetail:   PaymentDetail{},
			Price:           payment.Sum,
			PaymentDetailId: 14,
		}
		db.Save(&paymentHistory)
	}

	//if err := db.Exec("update restaurants " +
	//	"join restaurant_levels rl on rl.id = restaurants.restaurant_level_id " +
	//	"join stadia s on s.id = restaurants.stadium_id " +
	//	"join teams t on t.id = s.team_id " +
	//	"join users u on t.id = u.team_id " +
	//	"set u.coin = CASE " +
	//	"WHEN rl.weekly_price_type = 1 THEN u.coin - rl.weekly_price" +
	//	"  ELSE u.coin END," +
	//	" u.gem  = CASE WHEN rl.weekly_price_type = 0 THEN u.gem - rl.weekly_price" +
	//	"  ELSE u.gem END," +
	//	" last_paid = now() " +
	//	"where day(timediff(now(), restaurants.last_earned)) > 6;").Error; err != nil {
	//	return errors.InternalServerError(err.Error())
	//}

	return nil
}

func (team *Team) RemoveDoctor(uid int64) *errors.RestErr {
	var teamId uint
	var workerId int64

	if err := db.Select("team_id").Table("users").
		Where("id = ?", uid).
		Scan(&teamId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Select("doctor_id").Table("teams").
		Where("id = ?", teamId).
		Scan(&workerId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", teamId).
		Updates(map[string]interface{}{"doctor_id": nil}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	event := "fired doctor level " + strconv.FormatInt(workerId, 10)
	if err := team.SubmitTransferData(uid, 0, event); err != nil {
		return errors.InternalServerError("something went wrong setting transfer data")
	}

	return nil
}

func (team *Team) RemoveAssistantCoach(uid int64) *errors.RestErr {
	var teamId uint
	var workerId int64

	if err := db.Select("team_id").Table("users").
		Where("id = ?", uid).
		Scan(&teamId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Select("assistant_coach_id").Table("teams").
		Where("id = ?", teamId).
		Scan(&workerId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", teamId).
		Updates(map[string]interface{}{"assistant_coach_id": nil}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	event := "fired assistant coach level " + strconv.FormatInt(workerId, 10)
	if err := team.SubmitTransferData(uid, 0, event); err != nil {
		return errors.InternalServerError("something went wrong setting transfer data")
	}

	return nil
}

func (team *Team) RemoveFitnessCoach(uid int64) *errors.RestErr {
	var teamId uint
	var workerId int64
	if err := db.Select("team_id").Table("users").
		Where("id = ?", uid).
		Scan(&teamId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Select("fitness_coach_id").Table("teams").
		Where("id = ?", teamId).
		Scan(&workerId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", teamId).
		Updates(map[string]interface{}{"fitness_coach_id": nil}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	event := "fired fitness coach level " + strconv.FormatInt(workerId, 10)
	if err := team.SubmitTransferData(uid, 0, event); err != nil {
		return errors.InternalServerError("something went wrong setting transfer data")
	}

	return nil
}

func (team *Team) RemoveTalentFinder(uid int64) *errors.RestErr {
	var teamId uint
	var workerId int64

	if err := db.Select("team_id").Table("users").
		Where("id = ?", uid).
		Scan(&teamId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Select("talent_finder_id").Table("teams").
		Where("id = ?", teamId).
		Scan(&workerId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", teamId).
		Updates(map[string]interface{}{"talent_finder_id": nil}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	event := "fired talent finder level " + strconv.FormatInt(workerId, 10)
	if err := team.SubmitTransferData(uid, 0, event); err != nil {
		return errors.InternalServerError("something went wrong setting transfer data")
	}

	return nil
}

func (team *Team) RemoveTrainer(uid int64) *errors.RestErr {
	var teamId uint
	var workerId int64

	if err := db.Select("team_id").Table("users").
		Where("id = ?", uid).Scan(&teamId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Select("trainer_id").Table("teams").
		Where("id = ?", teamId).
		Scan(&workerId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", teamId).
		Updates(map[string]interface{}{"trainer_id": nil}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	event := "fired trainer level " + strconv.FormatInt(workerId, 10)
	if err := team.SubmitTransferData(uid, 0, event); err != nil {
		return errors.InternalServerError("something went wrong setting transfer data")
	}

	return nil
}

func (team *Team) ChangeTacticFormation(uid int64, tacticFormation string) *errors.RestErr {
	//select * from users join teams t on t.id = users.team_id where users.id = 1;
	if err := db.Select("*").
		Table("users").
		Joins("join teams t on t.id = users.team_id").
		Where("users.id = ?", uid).
		Find(&team).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NotFoundError("team not found")
		}
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("teams").
		Where("id = ?", team.ID).
		Updates(map[string]interface{}{"tactic_formation": tacticFormation}).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (team *Team) GetTacticFormation(uid int64) (string, *errors.RestErr) {
	//select * from users join teams t on t.id = users.team_id where users.id = 1;
	var tacticFormation string
	if err := db.Select("tactic_formation").
		Table("users").
		Joins("join teams t on t.id = users.team_id").
		Where("users.id = ?", uid).
		Find(&tacticFormation).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.NotFoundError("team not found")
		}
		return "", errors.InternalServerError("something went wrong")
	}

	return tacticFormation, nil
}

func (team *Team) UpgradeSection(uid int64, section string) *errors.RestErr {
	var userScore int64
	var stadiumId int64
	sectionName := section + "s"
	sectionLevel := section + "_level_id"
	sectionLevelIncrement := sectionLevel + " + ?"

	if err := db.Table("users").
		Select("score").
		Where("id = ?", uid).
		Find(&userScore).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("users").Select("s.id").
		Joins("join teams t on t.id = users.team_id").
		Joins("join stadia s on t.id = s.team_id").
		Where("users.id = ?", uid).Find(&stadiumId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table(sectionName).
		Where("stadium_id = ?", stadiumId).
		Updates(map[string]interface{}{sectionLevel: gorm.Expr(sectionLevelIncrement, 1)}).
		Error; err != nil {
		if strings.Contains(err.Error(), "Error 1452") {
			return errors.NotFoundError("reached the max level")
		}
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (team *Team) UpgradeStadium(uid int64, section string) *errors.RestErr {
	var userScore int64
	var stadiumId int64
	sectionName := section + "s"
	sectionLevel := section + "_level_id"
	sectionLevelIncrement := sectionLevel + " + ?"

	if err := db.Table("users").
		Select("score").
		Where("id = ?", uid).
		Find(&userScore).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("users").Select("s.id").
		Joins("join teams t on t.id = users.team_id").
		Joins("join stadia s on t.id = s.team_id").
		Where("users.id = ?", uid).Find(&stadiumId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if sectionName == "lights" {
		sectionLevel = "lights_level_id"
		sectionLevelIncrement = "lights_level_id + ?"
	}

	if sectionName == "stadiums" {
		sectionName = "stadia"
		sectionLevel = "stadium_level_id"
		sectionLevelIncrement = "stadium_level_id + ?"
		if err := db.Table(sectionName).
			Where("id = ?", stadiumId).
			Updates(map[string]interface{}{sectionLevel: gorm.Expr(sectionLevelIncrement, 1)}).
			Error; err != nil {
			if strings.Contains(err.Error(), "Error 1452") {
				return errors.NotFoundError("reached the max level")
			}
			return errors.InternalServerError("something went wrong")
		}
		return nil
	}

	if err := db.Table(sectionName).
		Where("stadium_id = ?", stadiumId).
		Updates(map[string]interface{}{sectionLevel: gorm.Expr(sectionLevelIncrement, 1)}).
		Error; err != nil {
		if strings.Contains(err.Error(), "Error 1452") {
			return errors.NotFoundError("reached the max level")
		}
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (team *Team) ChangeTeamCaptain(uid, captainId int64) *errors.RestErr {
	var teamId, capTeamId uint

	if err := db.Select("t.id").
		Table("users").
		Joins("join teams t on t.id = users.team_id").
		Where("users.id = ?", uid).Scan(&teamId).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Select("team_id").
		Table("players").
		Where("id = ?", captainId).Scan(&capTeamId).Error; err != nil {
		if strings.Contains(err.Error(), "converting NULL to uint is unsupported") {
			return errors.BadRequestError("captain not belong to your team")
		}
		return errors.InternalServerError("something went wrong")
	}

	if capTeamId != teamId {
		return errors.BadRequestError("captain not belong to your team")
	}

	if err := db.Table("teams").
		Where("id = ?", teamId).
		Updates(map[string]interface{}{"captain_id": captainId}).Error; err != nil {
		return errors.InternalServerError("something went wrong")

	}

	return nil
}

func (team *Team) ListOfIncome(uid int64) (map[string]int64, *errors.RestErr) {
	var teamId uint
	var stadium dto_only.Stadium
	income := make(map[string]int64)
	var parking int64
	var restaurant int64
	var shopping int64

	if err := db.Select("t.id").
		Table("users").
		Joins("join teams t on t.id = users.team_id").
		Where("users.id = ?", uid).Scan(&teamId).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if teamId == 0 {
		return nil, errors.BadRequestError("user does not have a team")
	}

	if err := db.Table("stadia").
		Where("team_id = ?", teamId).Scan(&stadium).Error; err != nil {
		if strings.Contains(err.Error(), "converting NULL to uint is unsupported") {
			return nil, errors.BadRequestError("captain not belong to your team")
		}
		return nil, errors.InternalServerError("something went wrong")
	}
	fmt.Println(stadium.ID)
	if stadium.ID == 0 {
		return nil, errors.BadRequestError("user does not have stadium")
	}

	if err := db.Select("daily_income_price").
		Table("parkings").
		Joins("join parking_levels pl on pl.id = parkings.parking_level_id").
		Where("parkings.stadium_id = ?", stadium.ID).Scan(&parking).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("daily_income_price").
		Table("restaurants").
		Joins("join restaurant_levels rl on rl.id = restaurants.restaurant_level_id").
		Where("restaurants.stadium_id = ?", stadium.ID).Scan(&restaurant).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("daily_income_price").
		Table("shoppings").
		Joins("join shopping_levels sl on sl.id = shoppings.shopping_level_id").
		Where("shoppings.stadium_id = ?", stadium.ID).Scan(&shopping).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	income["tickets income per game"] = stadium.TicketPrice * stadium.TotalCapacity
	income["restaurant daily income in coins"] = restaurant
	income["parking daily income in coin"] = parking
	income["shopping daily income in coin"] = shopping

	return income, nil
}

func (team *Team) ListOfOutcome(uid int64) (map[string]int64, *errors.RestErr) {
	var teamId, stadiumId uint
	outcome := make(map[string]int64)
	var assistantCoachSalary, DoctorSalary,
		TalentFinderSalary, TrainerSalary,
		FitnessCoachSalary, playersPrice int64
	var parking levels.ParkingLevel
	var restraunt levels.RestaurantLevel
	var shopping levels.ShoppingLevel
	var ground levels.GroundLevels
	var transportation levels.TransportationLevel
	if err := db.Select("t.id").
		Table("users").
		Joins("join teams t on t.id = users.team_id").
		Where("users.id = ?", uid).Scan(&teamId).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if teamId == 0 {
		return nil, errors.BadRequestError("user does not have a team")
	}

	if err := db.Select("ac.weekly_salary").
		Table("teams").
		Joins("join assistant_coaches ac on teams.assistant_coach_id = ac.id").
		Where("teams.id = ?", teamId).
		Scan(&assistantCoachSalary).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("d.weekly_salary").
		Table("teams").
		Joins("join doctors d on teams.doctor_id = d.id").
		Where("teams.id = ?", teamId).
		Scan(&DoctorSalary).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("tf.weekly_salary").
		Table("teams").
		Joins("join talent_finders tf on teams.talent_finder_id = tf.id").
		Where("teams.id = ?", teamId).
		Scan(&TalentFinderSalary).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("t.weekly_salary").
		Table("teams").
		Joins("join trainers t on teams.trainer_id = t.id").
		Where("teams.id = ?", teamId).
		Scan(&TrainerSalary).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("fc.weekly_salary").
		Table("teams").
		Joins("join fitness_coaches fc on teams.fitness_coach_id = fc.id").
		Where("teams.id = ?", teamId).
		Scan(&FitnessCoachSalary).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	db.Select("SUM(price)").
		Table("players").
		Where("team_id = ?", teamId).Scan(&playersPrice)

	if err := db.Select("id").
		Table("stadia").
		Where("team_id = ?", teamId).Scan(&stadiumId).Error; err != nil {
		if strings.Contains(err.Error(), "converting NULL to uint is unsupported") {
			return nil, errors.BadRequestError("captain not belong to your team")
		}
		return nil, errors.InternalServerError("something went wrong")
	}

	if stadiumId == 0 {
		return nil, errors.BadRequestError("user does not have stadium")
	}

	if err := db.Select("pl.weekly_price_type, pl.weekly_price").
		Table("parkings").
		Joins("join parking_levels pl on pl.id = parkings.parking_level_id").
		Where("parkings.stadium_id = ?", stadiumId).Scan(&parking).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("rl.weekly_price_type, rl.weekly_price").
		Table("restaurants").
		Joins("join restaurant_levels rl on rl.id = restaurants.restaurant_level_id").
		Where("restaurants.stadium_id = ?", stadiumId).Scan(&restraunt).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("sl.weekly_price_type, sl.weekly_price").
		Table("shoppings").
		Joins("join shopping_levels sl on sl.id = shoppings.shopping_level_id").
		Where("shoppings.stadium_id = ?", stadiumId).Scan(&shopping).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("gl.weekly_price_type, gl.weekly_price").
		Table("grounds").
		Joins("join ground_levels gl on gl.id = grounds.ground_level_id").
		Where("grounds.stadium_id = ?", stadiumId).Scan(&ground).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("tl.weekly_price_type, tl.weekly_price").
		Table("transportations").
		Joins("join transportation_levels tl on tl.id = transportations.transportation_level_id").
		Where("transportations.stadium_id = ?", stadiumId).Scan(&transportation).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	outcome["assistant coach salary in coins"] = assistantCoachSalary
	outcome["doctor salary in coins"] = DoctorSalary
	outcome["fitness coach salary in coins"] = FitnessCoachSalary
	outcome["trainer salary in coins"] = TrainerSalary
	outcome["talent finder salary in coins"] = TalentFinderSalary
	outcome["players salary in coins"] = playersPrice

	if restraunt.WeeklyPriceType == 0 {
		outcome["restaurant price in gem"] = restraunt.WeeklyPrice
	}
	if restraunt.WeeklyPriceType == 1 {
		outcome["restaurant price in coins"] = restraunt.WeeklyPrice
	}
	if parking.WeeklyPriceType == 0 {
		outcome["parking price in gem"] = parking.WeeklyPrice
	}
	if parking.WeeklyPriceType == 1 {
		outcome["parking price in coin"] = parking.WeeklyPrice
	}
	if shopping.WeeklyPriceType == 0 {
		outcome["shopping price in gem"] = shopping.WeeklyPrice
	}
	if shopping.WeeklyPriceType == 1 {
		outcome["shopping price in coin"] = shopping.WeeklyPrice
	}
	if ground.WeeklyPriceType == 0 {
		outcome["ground price in gem"] = ground.WeeklyPrice
	}
	if ground.WeeklyPriceType == 1 {
		outcome["ground price in coin"] = ground.WeeklyPrice
	}
	if transportation.WeeklyPriceType == 0 {
		outcome["transportation price in gem"] = transportation.WeeklyPrice
	}
	if transportation.WeeklyPriceType == 1 {
		outcome["transportation price in coin"] = transportation.WeeklyPrice
	}

	return outcome, nil
}

func (team *Team) IdToName(id int64) (string, *errors.RestErr) {
	var teamName string

	if err := db.Select("name").
		Table("teams").
		Where("id = ?", id).
		Scan(&teamName).Error; err != nil {
		return "", errors.InternalServerError("something went wrong")
	}

	return teamName, nil
}

func (team *Team) ChangeTicketPrice(uid, price int64) *errors.RestErr {
	var teamId int64

	if err := db.Select("t.id").
		Table("users").
		Joins("join teams t on t.id = users.team_id").
		Where("users.id = ?", uid).
		Scan(&teamId).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if teamId == 0 {
		return errors.BadRequestError("user does not have a team")
	}

	if err := db.Table("stadia").Where("team_id = ?", teamId).
		Updates(map[string]interface{}{"ticket_price": price}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (team *Team) SubmitTransferData(uid, playerId int64, event string) *errors.RestErr {

	transferData := TransferData{
		Model:         gorm.Model{},
		User:          User{},
		UserId:        uint(uid),
		Event:         event,
		Player:        Player{},
		PlayerId:      uint(playerId),
		SubmittedDate: time.Now().UTC(),
	}

	fmt.Println("INJAAAAAM ID", uid)

	if err := db.Create(&transferData).Error; err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}

func (team *Team) GetTransferData(uid int64) (map[string]interface{}, *errors.RestErr) {
	var transferData []TransferDataMap
	transferDataMap := make(map[string]interface{})

	if err := db.Table("transfer_data").Where("user_id = ?", uid).
		Find(&transferData).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	transferDataMap["transfer_data_map"] = transferData

	return transferDataMap, nil
}

func (team *Team) GetAllTactics() ([]dto_only.GetAllTactics, *errors.RestErr) {
	var tactics []dto_only.GetAllTactics

	if err := db.Table("tactics").Find(&tactics).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return tactics, nil
}

func (team *Team) IncreaseStadiumSeats(uid int64, section string) *errors.RestErr {
	var userScore int64
	var stadiumId int64
	sectionName := section + "s"
	sectionLevel := section + "_level_id"
	sectionLevelIncrement := sectionLevel + " + ?"

	if err := db.Table("users").
		Select("score").
		Where("id = ?", uid).
		Find(&userScore).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("users").Select("s.id").
		Joins("join teams t on t.id = users.team_id").
		Joins("join stadia s on t.id = s.team_id").
		Where("users.id = ?", uid).Find(&stadiumId).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if sectionName == "lights" {
		sectionLevel = "lights_level_id"
		sectionLevelIncrement = "lights_level_id + ?"
	}

	if sectionName == "stadiums" {
		sectionName = "stadia"
		sectionLevel = "stadium_level_id"
		sectionLevelIncrement = "stadium_level_id + ?"
		if err := db.Table(sectionName).
			Where("id = ?", stadiumId).
			Updates(map[string]interface{}{sectionLevel: gorm.Expr(sectionLevelIncrement, 1)}).
			Error; err != nil {
			if strings.Contains(err.Error(), "Error 1452") {
				return errors.NotFoundError("reached the max level")
			}
			return errors.InternalServerError("something went wrong")
		}
		return nil
	}

	if err := db.Table(sectionName).
		Where("stadium_id = ?", stadiumId).
		Updates(map[string]interface{}{sectionLevel: gorm.Expr(sectionLevelIncrement, 1)}).
		Error; err != nil {
		if strings.Contains(err.Error(), "Error 1452") {
			return errors.NotFoundError("reached the max level")
		}
		return errors.InternalServerError("something went wrong")
	}

	return nil
}
