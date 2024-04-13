package domain

import (
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Post     Post            `json:"post"`
	PostId   uint            `json:"post_id" gorm:"not null;"`
	Status   dto_only.Status `json:"status"`
	StatusId uint            `json:"status_id" gorm:"default:null;"`
	Name     string          `json:"name"`
	Energy   int             `json:"energy" gorm:"default:70"`
	Age      int             `json:"age"`
	Skill    int             `json:"skill"`
	Cheer    int             `json:"cheer"`
	Goal     int             `json:"goal"`
	Price    int64           `json:"price"`
	TeamID   uint            `json:"team_id" gorm:"default:null;"`
	Salary   int64           `json:"salary" gorm:"default:null;"`
	Power    PlayerPower     `json:"player_power" gorm:"constraint:OnDelete:CASCADE"`
}

type GetPlayers struct {
	gorm.Model
	Post     Post
	PostId   uint `json:"post_id" gorm:"not null;"`
	Status   dto_only.Status
	StatusId uint        `json:"status_id" gorm:"default:null;"`
	Name     string      `json:"name"`
	Energy   int         `json:"energy" gorm:"default:70"`
	Age      int         `json:"age"`
	Skill    int         `json:"skill"`
	Cheer    int         `json:"cheer"`
	Goal     int         `json:"goal"`
	Price    int64       `json:"price"`
	TeamID   uint        `json:"team_id" gorm:"default:null;"`
	Power    PlayerPower `json:"player_power" gorm:"constraint:OnDelete:CASCADE"`
}

type PlayerFilter struct {
	gorm.Model
	PostId   uint   `json:"post_id" gorm:"not null;"`
	StatusId uint   `json:"status_id" gorm:"default:null;"`
	Name     string `json:"name"`
	Energy   int    `json:"energy" gorm:"default:70"`
	Age      int    `json:"age"`
	Skill    int    `json:"skill"`
	Cheer    int    `json:"cheer"`
	Goal     int    `json:"goal"`
	Price    int64  `json:"price"`
	TeamID   uint   `json:"team_id" gorm:"default:null;"`
}
