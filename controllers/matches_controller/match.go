package matches_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/services/matches_service"
	"github.com/kasrashrz/Affogato/utils/errors"
	"net/http"
	"strconv"
)

func Create(ctx *gin.Context) {
	var match domain.Match
	if err := ctx.ShouldBind(&match); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	saveErr := matches_service.MatchServices.Create(match)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "created"})
	return
}

func MatchNotification(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	result, err := matches_service.MatchServices.MatchNotification(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": result})
	return

}

func UpdateResult(ctx *gin.Context) {
	var match domain.Match
	if err := ctx.ShouldBind(&match); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	saveErr := matches_service.MatchServices.UpdateResult(match)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func UsersMatches(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}
	matches, saveErr := matches_service.MatchServices.UsersMatches(id)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": matches})
	return
}

func SingleMatch(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}
	match, saveErr := matches_service.MatchServices.SingleMatch(id)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": match})
	return
}

func FriendlyMatch(ctx *gin.Context) {
	var match domain.Match
	if err := ctx.ShouldBind(&match); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	saveErr := matches_service.MatchServices.FriendlyMatch(match)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "created"})
	return
}

func JoinFriendlyMatch(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("match-id"), 10, 64)

	if idErr != nil {
		fmt.Println(idErr)
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}
	joinErr := matches_service.MatchServices.JoinFriendlyMatch(id)
	if joinErr != nil {
		ctx.JSON(joinErr.Status, joinErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "joined"})
	return
}

func GetFriendlyMatchesInvites(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}
	matches, saveErr := matches_service.MatchServices.GetFriendlyMatchesInvites(id)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": matches})
	return
}
