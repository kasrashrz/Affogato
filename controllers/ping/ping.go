package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/domain"
	"net/http"
	"time"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
	return
}

func DatabaseSetup(ctx *gin.Context) {
	dao := domain.Setup{}
	if err := dao.Setup(); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "setup done"})
	return
}

func ServerTime(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"response": time.Now()})
	return
}
