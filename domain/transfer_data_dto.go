package domain

import (
	"gorm.io/gorm"
	"time"
)

type TransferData struct {
	gorm.Model
	User          User      `json:"-"`
	UserId        uint      `json:"user_id"`
	Event         string    `json:"event"`
	Player        Player    `json:"-"`
	PlayerId      uint      `json:"player_id" gorm:"default:null"`
	SubmittedDate time.Time `json:"submitted_date"`
}

type TransferDataMap struct {
	Event         string    `json:"event"`
	PlayerId      uint      `json:"player_id" gorm:"default:null"`
	SubmittedDate time.Time `json:"submitted_date"`
}
