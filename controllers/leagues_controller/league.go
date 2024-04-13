package leagues_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/services/leagues_service"
	"github.com/kasrashrz/Affogato/utils/errors"
	"net/http"
	"strconv"
)

func Create(ctx *gin.Context) {
	var league domain.League
	if err := ctx.ShouldBind(&league); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	err := leagues_service.LeagueServices.Create(league)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "create"})
	return
}

func TestAPI(ctx *gin.Context) {
	err := leagues_service.LeagueServices.TestAPI()
	ctx.JSON(http.StatusOK, gin.H{"response": err})
	return
}

func LeaderBoard(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)

	if uidErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	teams, err := leagues_service.LeagueServices.LeaderBoard(uid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": teams})
	return
}

func TodayMatches(ctx *gin.Context) {
	userId, userIdErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if userIdErr != nil {
		err := errors.BadRequestError("invalid user-id format")
		ctx.JSON(err.Status, err)
		return
	}

	todayMatches, err := leagues_service.LeagueServices.TodayMatches(userId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": todayMatches})
	return

}
