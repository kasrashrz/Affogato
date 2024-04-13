package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/utils/jwt"
	"net/http"
)

type TokenRequest struct {
	Id int64 `json:"id"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user domain.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record, restErr := user.GetOne(request.Id)
	if restErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": restErr.Message})
		context.Abort()
		return
	}
	tokenString, err := jwt.GenerateJWT(record.Email, "no-secret-for-generation")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
