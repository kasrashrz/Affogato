package domain

import (
	checkErr "errors"
	"fmt"
	_ "github.com/kasrashrz/Affogato/datastore/mysql"
	"github.com/kasrashrz/Affogato/logger"
	errors "github.com/kasrashrz/Affogato/utils/errors"
	"gorm.io/gorm"
	"strings"
)

func (user *User) GetAll() ([]User, *errors.RestErr) {
	var users []User
	db.Table("users").Limit(100).Scan(&users)
	return users, nil
}

func (user *User) GetOne(id int64) (*User, *errors.RestErr) {
	var outputUser User
	if err := db.
		Preload("Team.Stadium").
		Where("id = ?", id).
		First(&outputUser).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFoundError("user does not exist")
		}
		fmt.Println(err)
		return nil, errors.InternalServerError("something went wrong")
	}

	return &outputUser, nil
}

func (user *User) Create() *errors.RestErr {

	//users_doman.DateCreated = date_time.GetNowDbFormat()

	user.Coin = 0
	user.Gem = 0
	user.Score = 1
	user.GroupScore = 0
	if err := db.Save(&user).Error; err != nil {
		if strings.ContainsAny(err.Error(), "Error 1062") {
			return errors.BadRequestError("username exists")
		}
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (user *User) CoinIncrement(coin int64, id int64) *errors.RestErr {

	if err := db.Table("users").
		Where("id = ?", id).
		Update("coin", gorm.Expr("coin + ?", coin)).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (user *User) CoinDecrement(coin, id int64) *errors.RestErr {

	var currentCoins int64
	db.Table("users").Select("coin").Where("id = ?", id).Scan(&currentCoins)

	if currentCoins < coin {
		return errors.BadRequestError("not enough coin")
	}

	if err := db.Table("users").
		Where("id = ?", id).
		Update("coin", gorm.Expr("coin - ?", coin)).
		Error; err != nil {
		return errors.InternalServerError("something went wrong during the coin decrement")
	}

	return nil

}

func (user *User) GemIncrement(gem int64, id int64) *errors.RestErr {

	if err := db.Table("users").
		Where("id = ?", id).
		Update("gem", gorm.Expr("gem + ?", gem)).
		Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (user *User) GemDecrement(gem, id int64) *errors.RestErr {

	var currentGems int64
	db.Table("users").Select("gem").Where("id = ?", id).Scan(&currentGems)

	if currentGems < gem {
		return errors.BadRequestError("not enough gem")
	}

	if err := db.Table("users").
		Where("id = ?", id).
		Update("gem", gorm.Expr("gem - ?", gem)).
		Error; err != nil {
		return errors.InternalServerError("something went wrong during the gem decrement")
	}

	return nil

}

func (user *User) UpdateUsername(username string, id uint) *errors.RestErr {

	err := db.Table("users").Where("id = ?", id).Update("username", username)
	if err.Error != nil {
		logger.Error("something went wrong", err.Error)
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (user *User) Delete(id int64) *errors.RestErr {

	if err := db.Delete(&User{}, id).Error; err != nil {
		if checkErr.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFoundError("user does not exist")
		}
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (user *User) AddFriend(userOneId, userTwoId int64) *errors.RestErr {

	if err := db.Exec("insert into user_friends values (?,?)", userOneId, userTwoId).Error; err != nil {
		if strings.Contains(err.Error(), "Error 1452") {
			return errors.NotFoundError("users not found")
		}
		if strings.Contains(err.Error(), "Error 1062") {
			return errors.BadRequestError("already friends")
		}
		return errors.InternalServerError("something went wrong")
	}
	if err := db.Exec("insert into user_friends values (?,?)", userTwoId, userOneId).Error; err != nil {
		if strings.Contains(err.Error(), "Error 1452") {
			return errors.NotFoundError("users not found")
		}
		if strings.Contains(err.Error(), "Error 1062") {
			return errors.BadRequestError("already friends")
		}
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (user *User) GetFriends(id int64) (*errors.RestErr, []map[string]interface{}) {
	var friendsList []map[string]interface{}
	if err := db.Select("u.username, u.id").Table("user_friends").
		Joins("join users u on u.id = user_friends.friend_id").
		Where("user_id = ?", id).
		Scan(&friendsList).Error; err != nil {
		return errors.InternalServerError("something went wrong"), nil
	}

	return nil, friendsList
}

func (user *User) ChangeAvatarFormation() *errors.RestErr {
	if err := db.Table("users").
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{"avatar_formation": user.AvatarFormation}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (user *User) ChangeToken(email, token, secret string) *errors.RestErr {
	if err := db.Table("users").
		Where("email = ?", email).
		Updates(map[string]interface{}{"token": token, "secret": secret}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (user *User) GetAvatarFormation(uid int64) (string, *errors.RestErr) {
	var avatarFormat string
	if err := db.Select("avatar_formation").Table("users").
		Where("id = ?", uid).
		Find(&avatarFormat).Error; err != nil {
		return "", errors.InternalServerError("something went wrong")
	}

	return avatarFormat, nil
}

func (user *User) GetUidFromToken(email string) (uint, *errors.RestErr) {
	var uid uint
	if err := db.Select("id").Table("users").
		Where("email = ?", email).Find(&uid).
		Error; err != nil {
		return 0, errors.InternalServerError("something went wrong")
	}
	return uid, nil
}

func (user *User) GetAllMissions() ([]Mission, *errors.RestErr) {
	var missions []Mission
	if err := db.Table("missions").
		Find(&missions).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return missions, nil
}

func (user *User) GetUserMissions(uid int64) ([]map[string]interface{}, *errors.RestErr) {
	var missions []map[string]interface{}
	if err := db.Select("missions.id, missions.name, missions.prize").
		Table("missions").
		Joins("join users_missions um on missions.id = um.mission_id").
		Where("user_id = ?", uid).
		Find(&missions).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return missions, nil
}
