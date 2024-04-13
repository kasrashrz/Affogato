package domain

import (
	"gorm.io/gorm"
)

type Cup struct {
	gorm.Model
	Name              string `json:"name"`
	MaxCapacity       int    `json:"max_capacity"`
	Invites           int    `json:"invites" gorm:"default:1"`
	EntryPricePerTeam int64  `json:"entry_price_per_team"`
	Type              bool   `json:"type" gorm:"default:0"`
}
