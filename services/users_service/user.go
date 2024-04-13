package users_service

import (
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/utils/errors"
	"github.com/kasrashrz/Affogato/utils/jwt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	UserServices userServiceInterface = &usersService{}
)

type usersService struct{}

type userServiceInterface interface {
	GetAll() ([]domain.User, *errors.RestErr)
	GetOne(id int64) (*domain.User, *errors.RestErr)
	CreateUser(user domain.User, secret, tokenString string) (*domain.User, *errors.RestErr)
	IncreaseUsersCoin(coin, id int64) *errors.RestErr
	DecreaseUsersCoin(coin, id int64) *errors.RestErr
	IncreaseUsersGem(coin, id int64) *errors.RestErr
	DecreaseUsersGem(coin, id int64) *errors.RestErr
	UpdateUsername(username string, id uint) (domain.User, *errors.RestErr)
	Delete(id int64) *errors.RestErr
	AddFriend(userOneId, userTwoId int64) *errors.RestErr
	GetFriends(id int64) (*errors.RestErr, []map[string]interface{})
	ChangeAvatarFormation(uid int64, avatarFormation string) *errors.RestErr
	GetAvatarFormation(uid int64) (string, *errors.RestErr)
	GetUidFromToken(token, secret string) (uint, *errors.RestErr)
	GetAllMissions() ([]domain.Mission, *errors.RestErr)
	GetUserMissions(uid int64) ([]map[string]interface{}, *errors.RestErr)
}

func (service *usersService) GetOne(id int64) (*domain.User, *errors.RestErr) {
	dao := domain.User{}
	user, err := dao.GetOne(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *usersService) GetAll() ([]domain.User, *errors.RestErr) {
	dao := domain.User{}

	return dao.GetAll()
}

func (service *usersService) CreateUser(user domain.User, secret, tokenString string) (*domain.User, *errors.RestErr) {
	user.Email = jwt.GetTokenData(tokenString, secret)
	if !strings.Contains(user.Email, "@") {
		return nil, errors.BadRequestError("wrong email")
	}
	if strings.ContainsAny(user.Username, "!@#$%^&*()") {
		return nil, errors.BadRequestError("wrong username")
	}
	if user.Username == "" {
		rand.Seed(time.Now().UnixNano())
		min := 100000
		max := 30000000
		user.Username = "guest" + strconv.Itoa(rand.Intn(max-min+1)+min)
	}
	if user.AvatarFormation == "" {
		user.AvatarFormation = "Lip1,Eye1,Nose1,EyeBrow1,SunGlass1,SkinColor1,HairStyle1,Mustache1,Shirt1,Collar1"
	}
	if err := user.Create(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (service *usersService) IncreaseUsersCoin(coin, id int64) *errors.RestErr {
	dao := domain.User{}
	return dao.CoinIncrement(coin, id)
}

func (service *usersService) DecreaseUsersCoin(coin, id int64) *errors.RestErr {
	dao := domain.User{}
	return dao.CoinDecrement(coin, id)
}

func (service *usersService) IncreaseUsersGem(coin, id int64) *errors.RestErr {
	dao := domain.User{}
	return dao.GemIncrement(coin, id)
}

func (service *usersService) DecreaseUsersGem(coin, id int64) *errors.RestErr {
	dao := domain.User{}
	return dao.GemDecrement(coin, id)
}

func (service *usersService) UpdateUsername(username string, id uint) (domain.User, *errors.RestErr) {
	if strings.ContainsAny(username, "!@#$%^&*()") {
		return domain.User{}, errors.BadRequestError("wrong username")
	}
	if len(username) == 0 {
		return domain.User{}, errors.BadRequestError("username could not be null")
	}
	dao := domain.User{}
	return dao, dao.UpdateUsername(username, id)
}

func (service *usersService) Delete(id int64) *errors.RestErr {
	dao := domain.User{}
	return dao.Delete(id)
}

func (service *usersService) AddFriend(userOneId, userTwoId int64) *errors.RestErr {
	dao := domain.User{}
	return dao.AddFriend(userOneId, userTwoId)
}

func (service *usersService) GetFriends(id int64) (*errors.RestErr, []map[string]interface{}) {
	dao := domain.User{}
	return dao.GetFriends(id)
}

func (service *usersService) ChangeAvatarFormation(uid int64, avatarFormation string) *errors.RestErr {
	dao := domain.User{}
	if strings.Contains(avatarFormation, "Lip") == false ||
		strings.Contains(avatarFormation, "Eye") == false ||
		strings.Contains(avatarFormation, "Nose") == false ||
		strings.Contains(avatarFormation, "EyeBrow") == false ||
		strings.Contains(avatarFormation, "SunGlass") == false ||
		strings.Contains(avatarFormation, "HairStyle") == false ||
		strings.Contains(avatarFormation, "Mustache") == false ||
		strings.Contains(avatarFormation, "Shirt") == false ||
		strings.Contains(avatarFormation, "Collar") == false ||
		strings.Contains(avatarFormation, "SkinColor") == false {
		return errors.BadRequestError("invalid avatar formation")
	}
	dao.ID = uint(uid)
	dao.AvatarFormation = avatarFormation
	return dao.ChangeAvatarFormation()
}

func (service *usersService) GetAvatarFormation(uid int64) (string, *errors.RestErr) {
	dao := domain.User{}
	return dao.GetAvatarFormation(uid)
}

func (service *usersService) GetUidFromToken(token, secret string) (uint, *errors.RestErr) {
	dao := domain.User{}
	email := jwt.GetTokenData(token, secret)
	return dao.GetUidFromToken(email)
}

func (service *usersService) GetAllMissions() ([]domain.Mission, *errors.RestErr) {
	dao := domain.User{}
	return dao.GetAllMissions()
}

func (service *usersService) GetUserMissions(uid int64) ([]map[string]interface{}, *errors.RestErr) {
	dao := domain.User{}
	return dao.GetUserMissions(uid)
}
