package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Friends         []*User `json:"-" gorm:"many2many:user_friends"`
	Username        string  `json:"username" gorm:"unique"`
	Email           string  `json:"email"`
	Coin            int64   `json:"coin" gorm:"default:500000;"`
	Gem             int64   `json:"gem" gorm:"default:200;"`
	Team            Team
	TeamId          uint   `json:"team_id" gorm:"default:null;"`
	Score           int64  `json:"score"`
	GroupScore      int64  `json:"group_score"`
	Token           string `json:"-"`
	Secret          string `json:"-"`
	AvatarFormation string `json:"avatar_formation"`
}

type UpdateUsername struct {
	Id       uint   `json:"id"`
	Username string `json:"username" gorm:"unique"`
}
