package players_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"github.com/kasrashrz/Affogato/services/players_service"
	"github.com/kasrashrz/Affogato/utils/errors"
	"log"
	"net/http"
	"strconv"
)

func GetAll(ctx *gin.Context) {
	player, getErr := players_service.PlayerServices.GetAll()
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": player})
	return
}

func GetOne(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	player, err := players_service.PlayerServices.GetOne(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": player})
	return
}

func Create(ctx *gin.Context) {
	var player domain.Player
	if err := ctx.ShouldBind(&player); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	saveErr := players_service.PlayerServices.Create(player)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "created"})
	return
}

func GetPlayersWithoutTeam(ctx *gin.Context) {
	player, err := players_service.PlayerServices.GetPlayersWithoutTeam()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": player})
	return
}

func GenerateRandomPlayers(ctx *gin.Context) {
	quantity, quantityError := strconv.Atoi(ctx.Query("quantity"))
	if quantityError != nil {
		err := errors.BadRequestError("invalid quantity format")
		ctx.JSON(err.Status, err)
		return
	}
	if err := players_service.PlayerServices.GenerateRandom(quantity); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}

func GoalIncrement(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	goal, goalErr := strconv.ParseInt(ctx.Query("goals"), 10, 64)

	if idErr != nil || goalErr != nil {
		err := errors.BadRequestError("invalid goal or id format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := players_service.PlayerServices.GoalIncrement(id, goal); err != nil {
		log.Println(err)
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func SetDefaultGoals(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := players_service.PlayerServices.SetDefaultGoals(id); err != nil {

		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func IncreaseAge(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	log.Println(idErr)
	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}
	if err := players_service.PlayerServices.IncreaseAge(id); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func RemoveOlds(ctx *gin.Context) {

	if err := players_service.PlayerServices.RemoveOlds(); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "removed"})
	return
}

func ChangeStatus(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	statusId, statusIdErr := strconv.ParseInt(ctx.Query("status_id"), 10, 64)

	if idErr != nil || statusIdErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}
	if err := players_service.PlayerServices.ChangeStatus(id, statusId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func Delete(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if idErr != nil {
		returnErr := errors.BadRequestError("invalid id format")
		ctx.JSON(returnErr.Status, returnErr)
		return
	}

	if err := players_service.PlayerServices.Delete(id); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "deleted"})
	return
}

func PlayerAvg(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	avg, err := players_service.PlayerServices.PlayerAvg(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"response": avg})
	return

}

func PowerList(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	list, err := players_service.PlayerServices.PowerList(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"response": list})
	return

}

func BodyBuilder(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	energy, energyErr := strconv.ParseInt(ctx.Query("energy"), 10, 64)

	if idErr != nil || energyErr != nil {
		err := errors.BadRequestError("invalid goal or id format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := players_service.PlayerServices.BodyBuilder(id, int(energy)); err != nil {
		log.Println(err)
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func RemovePlayerFromTeam(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid goal or id format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := players_service.PlayerServices.RemovePlayerFromTeam(id); err != nil {
		log.Println(err)
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "removed"})
	return
}

func SearchP2PWithFilter(ctx *gin.Context) {
	var filters dto_only.AutoGenerated
	if err := ctx.BindJSON(&filters); err != nil {
		fmt.Println(err)
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	players, saveErr := players_service.PlayerServices.SearchP2PWithFilter(filters.PlayerFilters)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": players})
	return
}

func SearchWithFilter(ctx *gin.Context) {
	var filters dto_only.AutoGenerated
	if err := ctx.BindJSON(&filters); err != nil {
		fmt.Println(err)
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	players, saveErr := players_service.PlayerServices.SearchP2PWithFilter(filters.PlayerFilters)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": players})
	return
}

func AddToBid(ctx *gin.Context) {
	var bid domain.Bid
	if err := ctx.BindJSON(&bid); err != nil {
		fmt.Println(err)
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	err := players_service.PlayerServices.AddToBid(bid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "inserted"})
	return
}

func GetAllBids(ctx *gin.Context) {
	bids, getErr := players_service.PlayerServices.GetAllBids()
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": bids})
	return
}

func GetOneBid(ctx *gin.Context) {
	bidId, bidIdErr := strconv.ParseInt(ctx.Query("bid-id"), 10, 64)
	if bidIdErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}
	bid, err := players_service.PlayerServices.GetOneBid(bidId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": bid})
	return
}

func IncreaseBid(ctx *gin.Context) {
	var bid domain.Bid
	if err := ctx.ShouldBind(&bid); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	err := players_service.PlayerServices.IncreaseBid(int64(bid.ID), int64(bid.BuyerID), bid.Price)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "increased"})
	return
}
