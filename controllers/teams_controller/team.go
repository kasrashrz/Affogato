package teams_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/services/teams_service"
	"github.com/kasrashrz/Affogato/utils/errors"
	"net/http"
	"strconv"
)

func GetAll(ctx *gin.Context) {
	teams, err := teams_service.TeamServices.GetAll()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": teams})
	return
}

func GetOne(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	team, err := teams_service.TeamServices.GetOne(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": team})
}

func Create(ctx *gin.Context) {
	var team domain.Team

	uid, _ := strconv.ParseInt(ctx.Query("uid"), 10, 64)

	if err := ctx.ShouldBind(&team); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	saveErr := teams_service.TeamServices.CreateTeam(uid, team)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "created"})
	return
}

func GetPlayers(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}
	players, err := teams_service.TeamServices.GetPlayers(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": players})
	return
}

func TeamPowerAvg(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	avg, err := teams_service.TeamServices.TeamPowerAvg(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": avg})
	return
}

func TeamEnergyAvg(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	avg, err := teams_service.TeamServices.TeamEnergyAvg(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": avg})
	return
}

func Delete(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.Delete(id); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "deleted"})
	return
}

func ChangeScore(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	score, scoreErr := strconv.ParseInt(ctx.Query("score"), 10, 64)
	if idErr != nil || scoreErr != nil {
		err := errors.BadRequestError("invalid id or score format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.ChangeScore(id, score); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func ChangeName(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	name := ctx.Query("name")
	if uidErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.ChangeName(uid, name); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func RandomGenerate(ctx *gin.Context) {
	quantity, quantityErr := strconv.ParseInt(ctx.Query("quantity"), 10, 64)

	if quantityErr != nil {
		err := errors.BadRequestError("invalid quantity format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.RandomGenerate(quantity); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "created"})
}

func TeamTactic(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	userId, userIdErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)

	if idErr != nil || userIdErr != nil {
		err := errors.BadRequestError("invalid id or uid format")
		ctx.JSON(err.Status, err)
		return
	}

	tactic, err := teams_service.TeamServices.TeamTactic(id, userId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reponse": tactic})
	return
}

func Workers(ctx *gin.Context) {
	userId, userIdErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if userIdErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	assistantCoach, trainer, fitnessCoach, doctor, talentFiner, err := teams_service.TeamServices.Workers(userId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"assistant_coach": assistantCoach,
		"trainer":         trainer,
		"fitness_coach":   fitnessCoach,
		"doctor":          doctor,
		"talent_finder":   talentFiner,
	})
	return
}

func BuyAssistantCoach(ctx *gin.Context) {
	userId, userIdErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	assistantCoachId, assistantCoachIdErr := strconv.ParseInt(ctx.Query("assistant-coach-id"), 10, 64)
	if userIdErr != nil || assistantCoachIdErr != nil {
		err := errors.BadRequestError("invalid id formats")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.BuyAssistantCoach(userId, assistantCoachId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func BuyDoctor(ctx *gin.Context) {
	userId, userIdErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	doctorId, doctorIdErr := strconv.ParseInt(ctx.Query("doctor-id"), 10, 64)
	if userIdErr != nil || doctorIdErr != nil {
		err := errors.BadRequestError("invalid id formats")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.BuyDoctor(userId, doctorId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func BuyFitnessCoach(ctx *gin.Context) {
	userId, userIdErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	fitnessCoachId, fitnessCoachIdErr := strconv.ParseInt(ctx.Query("fitness-coach-id"), 10, 64)
	if userIdErr != nil || fitnessCoachIdErr != nil {
		err := errors.BadRequestError("invalid id formats")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.BuyFitnessCoach(userId, fitnessCoachId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func BuyTalentFinder(ctx *gin.Context) {
	userId, userIdErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	talentFinderId, talentFinderIdErr := strconv.ParseInt(ctx.Query("talent-finder-id"), 10, 64)
	if userIdErr != nil || talentFinderIdErr != nil {
		err := errors.BadRequestError("invalid id formats")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.BuyTalentFinder(userId, talentFinderId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func BuyTrainer(ctx *gin.Context) {
	userId, userIdErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	trainerId, trainerIdErr := strconv.ParseInt(ctx.Query("trainer-id"), 10, 64)
	if userIdErr != nil || trainerIdErr != nil {
		err := errors.BadRequestError("invalid id formats")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.BuyTrainer(userId, trainerId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func GetPurchaseHistory(ctx *gin.Context) {
	userId, userIdErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if userIdErr != nil {
		err := errors.BadRequestError("invalid id formats")
		ctx.JSON(err.Status, err)
		return
	}

	purchases, err := teams_service.TeamServices.GetPurchaseHistory(userId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": purchases})
	return
}

func GetAllDoctors(ctx *gin.Context) {
	doctors, err := teams_service.TeamServices.GetAllDoctors()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": doctors})
	return
}

func GetAllTrainers(ctx *gin.Context) {
	trainers, err := teams_service.TeamServices.GetAllTrainers()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": trainers})
	return
}

func GetAllAssistantCoaches(ctx *gin.Context) {
	assistantCoaches, err := teams_service.TeamServices.GetAllAssistantCoaches()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": assistantCoaches})
	return
}

func GetAllFitnessCoaches(ctx *gin.Context) {
	fitnessCoaches, err := teams_service.TeamServices.GetAllFitnessCoaches()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": fitnessCoaches})
	return
}

func GetAllTalentFinders(ctx *gin.Context) {
	talentFinders, err := teams_service.TeamServices.GetAllTalentFinders()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": talentFinders})
	return
}

func ChangeTactic(ctx *gin.Context) {
	uid, idErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	tacticId, tacticIdErr := strconv.ParseInt(ctx.Query("tactic-id"), 10, 64)

	if idErr != nil || tacticIdErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.ChangeTactic(uid, tacticId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func GetStadiumData(ctx *gin.Context) {
	uid, idErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	stadium, err := teams_service.TeamServices.GetStadiumData(uid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": stadium})
	return
}

func GetStadiumLevels(ctx *gin.Context) {

	stadiumLevels, err := teams_service.TeamServices.GetStadiumLevels()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": stadiumLevels})
	return
}

func RemoveDoctor(ctx *gin.Context) {
	_, _ = strconv.ParseInt(ctx.Query("uid"), 10, 64)
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	err := teams_service.TeamServices.RemoveDoctor(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "removed"})
	return
}

func RemoveAssistantCoach(ctx *gin.Context) {
	_, _ = strconv.ParseInt(ctx.Query("uid"), 10, 64)
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	err := teams_service.TeamServices.RemoveAssistantCoach(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "removed"})
	return
}

func RemoveFitnessCoach(ctx *gin.Context) {
	_, _ = strconv.ParseInt(ctx.Query("uid"), 10, 64)
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	err := teams_service.TeamServices.RemoveFitnessCoach(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "removed"})
	return
}

func RemoveTalentFinder(ctx *gin.Context) {
	_, _ = strconv.ParseInt(ctx.Query("uid"), 10, 64)
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	err := teams_service.TeamServices.RemoveTalentFinder(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "removed"})
	return
}

func RemoveTrainer(ctx *gin.Context) {
	_, _ = strconv.ParseInt(ctx.Query("uid"), 10, 64)
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	err := teams_service.TeamServices.RemoveTrainer(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "removed"})
	return
}

func ChangeTacticFormation(ctx *gin.Context) {
	var team domain.Team
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)

	if uidErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := ctx.BindJSON(&team); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	if err := teams_service.TeamServices.ChangeTacticFormation(uid, team.TacticFormation); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func GetTacticFormation(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)

	if uidErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	tactic, err := teams_service.TeamServices.GetTacticFormation(uid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": tactic})
	return
}

func UpgradeSection(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	sectionName := ctx.Query("section-name")
	if uidErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.UpgradeSection(uid, sectionName); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func UpgradeStadium(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	sectionName := ctx.Query("section-name")
	if uidErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.UpgradeStadium(uid, sectionName); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func ChangeTeamCaptain(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	capId, capIdErr := strconv.ParseInt(ctx.Query("cap_id"), 10, 64)
	if uidErr != nil || capIdErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := teams_service.TeamServices.ChangeTeamCaptain(uid, capId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func ListOfIncome(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if uidErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	income, err := teams_service.TeamServices.ListOfIncome(uid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": income})
	return
}

func ListOfOutcome(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if uidErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	outcome, err := teams_service.TeamServices.ListOfOutcome(uid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": outcome})
	return
}

func IdToName(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if idErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	name, err := teams_service.TeamServices.IdToName(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": name})
	return
}

func ChangeTicketPrice(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	price, priceErr := strconv.ParseInt(ctx.Query("price"), 10, 64)
	if uidErr != nil || priceErr != nil {
		err := errors.BadRequestError("invalid uid or price format")
		ctx.JSON(err.Status, err)
		return
	}

	err := teams_service.TeamServices.ChangeTicketPrice(uid, price)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func GetTransferData(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if uidErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	transferDatas, err := teams_service.TeamServices.GetTransferData(uid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": transferDatas})
	return
}

func GetAllTactics(ctx *gin.Context) {
	tactics, err := teams_service.TeamServices.GetAllTactics()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": tactics})
	return
}
