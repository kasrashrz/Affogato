package domain

import (
	"gorm.io/gorm"
)

type League struct {
	gorm.Model
	LeagueName string `json:"league_name"`
	LeagueRate int    `json:"league_rate"`

	//EndingTime time.Time `json:"ending_time"`
}
