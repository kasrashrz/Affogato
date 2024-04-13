package cups_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/services/cups_service"
	"github.com/kasrashrz/Affogato/utils/errors"
	"net/http"
	"strconv"
)

func Create(ctx *gin.Context) {
	var cup domain.Cup

	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)

	if uidErr != nil {
		restError := errors.BadRequestError("invalid uid format")
		ctx.JSON(restError.Status, restError)
		return
	}

	if err := ctx.ShouldBind(&cup); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	saveErr := cups_service.CupServices.Create(cup, uid)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "created"})
	return
}

func GetAll(ctx *gin.Context) {
	_, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if uidErr != nil {
		restError := errors.BadRequestError("invalid uid format")
		ctx.JSON(restError.Status, restError)
		return
	}

	cups, saveErr := cups_service.CupServices.GetAll()
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": cups})
	return
}

func GetOne(ctx *gin.Context) {
	_, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	cupId, cupIdErr := strconv.ParseInt(ctx.Query("cup-id"), 10, 64)
	if uidErr != nil || cupIdErr != nil {
		restError := errors.BadRequestError("invalid uid format")
		ctx.JSON(restError.Status, restError)
		return
	}

	cup, cupErr := cups_service.CupServices.GetOne(cupId)
	if cupErr != nil {
		ctx.JSON(cupErr.Status, cupErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": cup})
	return
}

func Join(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	cupId, cupIdErr := strconv.ParseInt(ctx.Query("cup-id"), 10, 64)

	if uidErr != nil || cupIdErr != nil {
		restError := errors.BadRequestError("invalid uid format")
		ctx.JSON(restError.Status, restError)
		return
	}

	joinErr := cups_service.CupServices.Join(uid, cupId)
	if joinErr != nil {
		ctx.JSON(joinErr.Status, joinErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "joined"})
	return
}

func GetCupMatches(ctx *gin.Context) {
	_, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	cupId, cupIdErr := strconv.ParseInt(ctx.Query("cup-id"), 10, 64)

	if uidErr != nil || cupIdErr != nil {
		restError := errors.BadRequestError("invalid uid format")
		ctx.JSON(restError.Status, restError)
		return
	}

	matches, joinErr := cups_service.CupServices.GetCupMatches(cupId)
	if joinErr != nil {
		ctx.JSON(joinErr.Status, joinErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": matches})
	return
}

func UserLeagueAndCupData(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)

	if uidErr != nil {
		restError := errors.BadRequestError("invalid uid format")
		ctx.JSON(restError.Status, restError)
		return
	}

	data, joinErr := cups_service.CupServices.UserLeagueAndCupData(uid)

	if joinErr != nil {
		ctx.JSON(joinErr.Status, joinErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": data})
	return
}
