package markets_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/services/market_service"
	"github.com/kasrashrz/Affogato/utils/errors"
	"net/http"
	"strconv"
)

func AddPlayer(ctx *gin.Context) {
	var market domain.Market
	if err := ctx.BindJSON(&market); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}

	err := market_service.MarketServices.AddPlayer(market)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "create"})
	return
}

func BuyPlayer(ctx *gin.Context) {
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	marketId, marketIdErr := strconv.ParseInt(ctx.Query("market-id"), 10, 64)

	if uidErr != nil || marketIdErr != nil {
		err := errors.BadRequestError("invalid uid or market-id format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := market_service.MarketServices.BuyPlayer(uid, marketId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return

}

func AllMarketChoices(ctx *gin.Context) {
	start, startErr := strconv.ParseInt(ctx.Query("start"), 10, 64)
	count, countErr := strconv.ParseInt(ctx.Query("count"), 10, 64)

	if startErr != nil || countErr != nil {
		err := errors.BadRequestError("invalid start or count format")
		ctx.JSON(err.Status, err)
		return
	}

	result, err := market_service.MarketServices.AllMarketChoices(start, count)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": result})
}
