package users_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/services/users_service"
	"github.com/kasrashrz/Affogato/utils/errors"
	"net/http"
	"strconv"
)

func GetAll(ctx *gin.Context) {
	user, err := users_service.UserServices.GetAll()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": user})
}

func GetOne(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	result, err := users_service.UserServices.GetOne(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": result})
	return

}

func Create(ctx *gin.Context) {
	var user domain.User
	tokenString := ctx.GetHeader("Authorization")
	secret := ctx.GetHeader("Secret")

	if err := ctx.ShouldBind(&user); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	result, saveErr := users_service.UserServices.CreateUser(user, secret, tokenString)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusOK, result)
	return
}

func IncreaseCoin(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	coin, coinErr := strconv.ParseInt(ctx.Query("coin"), 10, 64)
	if coinErr != nil || idErr != nil {
		err := errors.BadRequestError("invalid id or coin format")
		ctx.JSON(err.Status, err)
		return
	}
	err := users_service.UserServices.IncreaseUsersCoin(coin, id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func DecreaseCoin(ctx *gin.Context) {

	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	coin, coinErr := strconv.ParseInt(ctx.Query("coin"), 10, 64)
	if coinErr != nil || idErr != nil {
		err := errors.BadRequestError("invalid id or coin format")
		ctx.JSON(err.Status, err)
		return
	}
	err := users_service.UserServices.DecreaseUsersCoin(coin, id)
	if err != nil {
		if err.Status != http.StatusInternalServerError {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err.Status == http.StatusInternalServerError {
			ctx.JSON(http.StatusInternalServerError, gin.H{"response": "something went wrong"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func IncreaseGem(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	gem, gemErr := strconv.ParseInt(ctx.Query("gem"), 10, 64)
	if gemErr != nil || idErr != nil {
		ctx.JSON(http.StatusBadRequest, errors.BadRequestError("invalid id or gem"))
		return
	}
	err := users_service.UserServices.IncreaseUsersGem(gem, id)

	if err != nil {
		if err.Status != http.StatusInternalServerError {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err.Status == http.StatusInternalServerError {
			ctx.JSON(http.StatusInternalServerError, gin.H{"response": "something went wrong"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func DecreaseGem(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	gem, gemErr := strconv.ParseInt(ctx.Query("gem"), 10, 64)
	if gemErr != nil || idErr != nil {
		ctx.JSON(http.StatusBadRequest, errors.BadRequestError("invalid id or gem"))
		return
	}
	err := users_service.UserServices.DecreaseUsersGem(gem, id)

	if err != nil {
		if err.Status != http.StatusInternalServerError {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err.Status == http.StatusInternalServerError {
			ctx.JSON(http.StatusInternalServerError, gin.H{"response": "something went wrong"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func UpdateUsername(ctx *gin.Context) {
	var user domain.UpdateUsername
	if err := ctx.ShouldBind(&user); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	_, err := users_service.UserServices.UpdateUsername(user.Username, user.Id)
	if err != nil {
		if err.Status != http.StatusInternalServerError {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		if err.Status == http.StatusInternalServerError {
			ctx.JSON(http.StatusInternalServerError, gin.H{"response": "something went wrong"})
			return
		}
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

	if err := users_service.UserServices.Delete(id); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "deleted"})
	return
}

func AddFriend(ctx *gin.Context) {

	userOneId, userOneIdErr := strconv.ParseInt(ctx.Query("uid-1"), 10, 64)
	userTwoId, userTwoIdErr := strconv.ParseInt(ctx.Query("uid-2"), 10, 64)

	if userOneIdErr != nil || userTwoIdErr != nil {
		err := errors.BadRequestError("invalid user id format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := users_service.UserServices.AddFriend(userOneId, userTwoId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "friend added"})
	return
}

func GetFriends(ctx *gin.Context) {

	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)

	if idErr != nil {
		err := errors.BadRequestError("invalid user id format")
		ctx.JSON(err.Status, err)
		return
	}

	err, usersFriends := users_service.UserServices.GetFriends(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": usersFriends})
	return
}

func ChangeAvatarFormat(ctx *gin.Context) {
	var user domain.User
	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if uidErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	if err := ctx.BindJSON(&user); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	if err := users_service.UserServices.ChangeAvatarFormation(uid, user.AvatarFormation); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": "updated"})
	return
}

func GetAvatarFormation(ctx *gin.Context) {

	uid, uidErr := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if uidErr != nil {
		err := errors.BadRequestError("invalid uid format")
		ctx.JSON(err.Status, err)
		return
	}

	avatarFormation, err := users_service.UserServices.GetAvatarFormation(uid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": avatarFormation})
	return
}

func GetUidFromEmail(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	secret := ctx.GetHeader("Secret")
	uid, err := users_service.UserServices.GetUidFromToken(tokenString, secret)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": uid})
	return
}

func GetUserMissions(ctx *gin.Context) {
	id, idErr := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if idErr != nil {
		err := errors.BadRequestError("invalid id format")
		ctx.JSON(err.Status, err)
		return
	}

	result, err := users_service.UserServices.GetUserMissions(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": result})
	return

}

func GetAllMissions(ctx *gin.Context) {
	result, err := users_service.UserServices.GetAllMissions()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": result})
	return

}
