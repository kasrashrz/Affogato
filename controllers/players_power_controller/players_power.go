package players_power_controller

import (
	"github.com/gin-gonic/gin"
	players_service "github.com/kasrashrz/Affogato/services/players_power_service"
	"github.com/kasrashrz/Affogato/utils/errors"
	"log"
	"net/http"
	"strconv"
)

func PracticePlayer(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	paramName := ctx.Query("param-name")

	if idErr != nil {
		err := errors.BadRequestError("invalid id, key or value format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := players_service.PlayerPowerServices.PracticePlayer(id, paramName); err != nil {
		log.Println(err)
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func SetDefaultPowers(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if idErr != nil {
		returnErr := errors.BadRequestError("invalid id format")
		ctx.JSON(returnErr.Status, returnErr)
		return
	}

	if err := players_service.PlayerPowerServices.SetDefaultPowers(id); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}
