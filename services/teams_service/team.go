package teams_service

import (
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"github.com/kasrashrz/Affogato/utils/errors"
	"math/rand"
	"strconv"
)

var (
	TeamServices teamServiceInterface = &teamsService{}
)

type teamsService struct{}

type teamServiceInterface interface {
	CreateTeam(uid int64, team domain.Team) *errors.RestErr
	GetAll() ([]domain.GetAllTeams, *errors.RestErr)
	GetOne(id int64) (*domain.Team, *errors.RestErr)
	GetPlayers(id int64) ([]domain.Player, *errors.RestErr)
	TeamPowerAvg(id int64) (int64, *errors.RestErr)
	TeamEnergyAvg(id int64) (int64, *errors.RestErr)
	Delete(id int64) *errors.RestErr
	ChangeScore(id, score int64) *errors.RestErr
	ChangeName(uid int64, name string) *errors.RestErr
	RandomGenerate(quantity int64) *errors.RestErr
	TeamTactic(id, userId int64) (map[string]string, *errors.RestErr)
	Workers(userId int64) (*dto_only.AssistantCoach, *dto_only.Trainer,
		*dto_only.FitnessCoach, *dto_only.Doctor,
		*dto_only.TalentFinder, *errors.RestErr)
	BuyAssistantCoach(userId int64, assistantCoachId int64) *errors.RestErr
	BuyDoctor(userId int64, doctorId int64) *errors.RestErr
	BuyFitnessCoach(userId int64, fitnessCoachId int64) *errors.RestErr
	BuyTalentFinder(userId int64, talentFinderId int64) *errors.RestErr
	BuyTrainer(userId int64, trainerId int64) *errors.RestErr
	GetPurchaseHistory(userId int64) (*[]domain.PaymentHistory, *errors.RestErr)
	GetAllDoctors() (*[]dto_only.Doctor, *errors.RestErr)
	GetAllTrainers() (*[]dto_only.Trainer, *errors.RestErr)
	GetAllAssistantCoaches() (*[]dto_only.AssistantCoach, *errors.RestErr)
	GetAllFitnessCoaches() (*[]dto_only.FitnessCoach, *errors.RestErr)
	GetAllTalentFinders() (*[]dto_only.TalentFinder, *errors.RestErr)
	ChangeTactic(uid, tacticId int64) *errors.RestErr
	GetStadiumData(uid int64) (map[string]interface{}, *errors.RestErr)
	GetStadiumLevels() (map[string]interface{}, *errors.RestErr)
	RemoveDoctor(id int64) *errors.RestErr
	RemoveAssistantCoach(id int64) *errors.RestErr
	RemoveFitnessCoach(id int64) *errors.RestErr
	RemoveTalentFinder(id int64) *errors.RestErr
	RemoveTrainer(id int64) *errors.RestErr
	ChangeTacticFormation(uid int64, tacticFormation string) *errors.RestErr
	GetTacticFormation(uid int64) (string, *errors.RestErr)
	UpgradeSection(uid int64, section string) *errors.RestErr
	UpgradeStadium(uid int64, section string) *errors.RestErr
	ChangeTeamCaptain(uid, captainId int64) *errors.RestErr
	ListOfIncome(uid int64) (map[string]int64, *errors.RestErr)
	ListOfOutcome(uid int64) (map[string]int64, *errors.RestErr)
	IdToName(id int64) (string, *errors.RestErr)
	ChangeTicketPrice(uid, price int64) *errors.RestErr
	SubmitTransferData(uid, playerId int64, event string) *errors.RestErr
	GetTransferData(uid int64) (map[string]interface{}, *errors.RestErr)
	GetAllTactics() ([]dto_only.GetAllTactics, *errors.RestErr)
}

func (service *teamsService) CreateTeam(uid int64, team domain.Team) *errors.RestErr {
	dao := domain.Team{}
	if team.Name == "" {
		team.Name = "Team" + strconv.Itoa(rand.Intn(999999-100000+1)+100000)
	}
	dao.TacticFormation = "12345678MDF"
	dao.Name = team.Name
	dao.LeagueId = team.LeagueId
	dao.TacticId = team.TacticId
	return dao.Create(uid)
}

func (service *teamsService) GetAll() ([]domain.GetAllTeams, *errors.RestErr) {
	dao := domain.GetAllTeams{}
	teams, err := dao.GetAll()
	if err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return teams, nil
}

func (service *teamsService) GetOne(id int64) (*domain.Team, *errors.RestErr) {
	dao := domain.GetAllTeams{}
	team, err := dao.GetOne(id)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (service *teamsService) GetPlayers(id int64) ([]domain.Player, *errors.RestErr) {
	dao := domain.Team{}
	return dao.FindPlayers(id)
}

func (service *teamsService) TeamPowerAvg(id int64) (avg int64, err *errors.RestErr) {
	dao := domain.Team{}
	return dao.TeamPowerAvg(id)
}

func (service *teamsService) TeamEnergyAvg(id int64) (avg int64, err *errors.RestErr) {
	dao := domain.Team{}
	return dao.TeamEnergyAvg(id)
}

func (service *teamsService) Delete(id int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.Delete(id)
}

func (service *teamsService) ChangeScore(id, score int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.ChangeScore(id, score)
}

func (service *teamsService) ChangeName(uid int64, name string) *errors.RestErr {
	dao := domain.Team{}
	return dao.ChangeName(uid, name)
}

func (service *teamsService) RandomGenerate(quantity int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.RandomGenerate(quantity)
}

func (service *teamsService) TeamTactic(id, userId int64) (map[string]string, *errors.RestErr) {
	dao := domain.Team{}
	tactic, err := dao.TeamTactic(id, userId)
	return tactic, err
}

func (service *teamsService) Workers(userId int64) (*dto_only.AssistantCoach, *dto_only.Trainer,
	*dto_only.FitnessCoach, *dto_only.Doctor,
	*dto_only.TalentFinder, *errors.RestErr) {
	dao := domain.Team{}
	assistantCoach, trainer, fitnessCoach, doctor, talentFinder, err := dao.Workers(userId)

	if assistantCoach.ID == 0 {
		assistantCoach = nil
	}
	if trainer.ID == 0 {
		trainer = nil
	}
	if fitnessCoach.ID == 0 {
		fitnessCoach = nil
	}
	if doctor.ID == 0 {
		doctor = nil
	}
	if talentFinder.ID == 0 {
		talentFinder = nil
	}

	return assistantCoach, trainer, fitnessCoach, doctor, talentFinder, err
}

func (service *teamsService) BuyAssistantCoach(userId int64, assistantCoachId int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.BuyAssistantCoach(userId, assistantCoachId)
}

func (service *teamsService) BuyDoctor(userId int64, doctorId int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.BuyDoctor(userId, doctorId)
}

func (service *teamsService) BuyFitnessCoach(userId int64, fitnessCoachId int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.BuyFitnessCoach(userId, fitnessCoachId)
}

func (service *teamsService) BuyTalentFinder(userId int64, talentFinderId int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.BuyTalentFinder(userId, talentFinderId)
}

func (service *teamsService) BuyTrainer(userId int64, trainerId int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.BuyTrainer(userId, trainerId)
}

func (service *teamsService) GetPurchaseHistory(userId int64) (*[]domain.PaymentHistory, *errors.RestErr) {
	dao := domain.Team{}
	purchases, err := dao.GetPurchaseHistory(userId)
	if err != nil {
		return nil, err
	}
	return purchases, nil
}

func (service *teamsService) GetAllDoctors() (*[]dto_only.Doctor, *errors.RestErr) {
	dao := domain.Team{}
	return dao.GetAllDoctors()
}

func (service *teamsService) GetAllTrainers() (*[]dto_only.Trainer, *errors.RestErr) {
	dao := domain.Team{}
	return dao.GetAllTrainers()
}

func (service *teamsService) GetAllAssistantCoaches() (*[]dto_only.AssistantCoach, *errors.RestErr) {
	dao := domain.Team{}
	return dao.GetAllAssistantCoaches()
}

func (service *teamsService) GetAllFitnessCoaches() (*[]dto_only.FitnessCoach, *errors.RestErr) {
	dao := domain.Team{}
	return dao.GetAllFitnessCoaches()
}

func (service *teamsService) GetAllTalentFinders() (*[]dto_only.TalentFinder, *errors.RestErr) {
	dao := domain.Team{}
	return dao.GetAllTalentFinders()
}

func (service *teamsService) ChangeTactic(uid, tacticId int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.ChangeTactic(uid, tacticId)
}

func (service *teamsService) GetStadiumData(uid int64) (map[string]interface{}, *errors.RestErr) {
	dao := domain.Team{}
	return dao.GetStadiumData(uid)
}

func (service *teamsService) GetStadiumLevels() (map[string]interface{}, *errors.RestErr) {
	dao := domain.Team{}
	return dao.GetStadiumLevels()
}

func (service *teamsService) RemoveDoctor(id int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.RemoveDoctor(id)
}

func (service *teamsService) RemoveAssistantCoach(id int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.RemoveAssistantCoach(id)
}

func (service *teamsService) RemoveFitnessCoach(id int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.RemoveFitnessCoach(id)
}

func (service *teamsService) RemoveTalentFinder(id int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.RemoveTalentFinder(id)
}

func (service *teamsService) RemoveTrainer(id int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.RemoveTrainer(id)
}

func (service *teamsService) ChangeTacticFormation(uid int64, tacticFormation string) *errors.RestErr {
	dao := domain.Team{}

	return dao.ChangeTacticFormation(uid, tacticFormation)
}

func (service *teamsService) GetTacticFormation(uid int64) (string, *errors.RestErr) {
	dao := domain.Team{}
	return dao.GetTacticFormation(uid)
}

func (service *teamsService) UpgradeSection(uid int64, section string) *errors.RestErr {
	dao := domain.Team{}
	return dao.UpgradeSection(uid, section)
}

func (service *teamsService) UpgradeStadium(uid int64, section string) *errors.RestErr {
	dao := domain.Team{}
	return dao.UpgradeStadium(uid, section)
}

func (service *teamsService) ChangeTeamCaptain(uid, captainId int64) *errors.RestErr {
	dao := domain.Team{}
	return dao.ChangeTeamCaptain(uid, captainId)
}

func (service *teamsService) ListOfIncome(uid int64) (map[string]int64, *errors.RestErr) {
	dao := domain.Team{}
	return dao.ListOfIncome(uid)
}

func (service *teamsService) ListOfOutcome(uid int64) (map[string]int64, *errors.RestErr) {
	dao := domain.Team{}
	return dao.ListOfOutcome(uid)
}

func (service *teamsService) IdToName(id int64) (string, *errors.RestErr) {
	dao := domain.Team{}

	return dao.IdToName(id)
}

func (service *teamsService) ChangeTicketPrice(uid, price int64) *errors.RestErr {
	dao := domain.Team{}
	if price < 500 {
		return errors.BadRequestError("not enough price for tickets")
	}
	if price > 2000 {
		return errors.BadRequestError("too much price for tickets")
	}
	return dao.ChangeTicketPrice(uid, price)
}

func (service *teamsService) SubmitTransferData(uid, playerId int64, event string) *errors.RestErr {
	dao := domain.Team{}
	return dao.SubmitTransferData(uid, playerId, event)
}

func (service *teamsService) GetTransferData(uid int64) (map[string]interface{}, *errors.RestErr) {
	dao := domain.Team{}
	return dao.GetTransferData(uid)
}

func (service *teamsService) GetAllTactics() ([]dto_only.GetAllTactics, *errors.RestErr) {
	dao := domain.Team{}
	return dao.GetAllTactics()
}
