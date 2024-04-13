package app

import (
	"fmt"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/configs"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/logger"
	"github.com/kasrashrz/Affogato/utils/errors"
	"github.com/kasrashrz/Affogato/utils/jwt"
	"net/http"
	"strconv"
	"time"
)

var (
	router = gin.Default()
)

func testResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusRequestTimeout, gin.H{"response": "timeout"})
}

func StartServer() {
	router.Use(timeoutMiddleware())
	//router.Use(loginRegister())
	mapUrls()
	db := domain.SetupModels() // Create models and set the first one manual
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})
	port := fmt.Sprintf(":%d", configs.ReadConfig().Server.Port)
	router.Run(port)
	logger.Info("Server started")
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(2000*time.Millisecond),
		timeout.WithHandler(func(ctx *gin.Context) {
			ctx.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}

func loginRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		//secret := ctx.GetHeader("Secret")

		isRegisteredStr := ctx.GetHeader("Is_Registered")
		isRegisteredInt, isRegisteredStrErr := strconv.ParseInt(isRegisteredStr, 10, 64)

		if isRegisteredStrErr != nil {
			err := errors.BadRequestError("invalid is registered format")
			ctx.AbortWithStatusJSON(err.Status, err)
			return
		}

		if isRegisteredInt == 0 {
			err := errors.Unauthorized("user have not registered yet")
			ctx.AbortWithStatusJSON(err.Status, err)
			return
		}

		if tokenString == "" && ctx.Request.URL.Path != "/login" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}

		if ctx.Request.URL.Path != "/login" {
			err := jwt.ValidateToken(tokenString, "affogato")
			if err != nil {
				ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
