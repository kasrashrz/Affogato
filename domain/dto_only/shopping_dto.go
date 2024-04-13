package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
	"time"
)

type Shopping struct {
	gorm.Model
	ShoppingLevel   levels.ShoppingLevel `json:"shopping_level"`
	ShoppingLevelId uint                 `json:"shopping_level_id" gorm:"default:null"`
	StadiumID       uint                 `json:"stadium_id" gorm:"default:null;"`
	LastPaid        time.Time            `json:"last_paid"`
	LastEarned      time.Time            `json:"last_earned"`
}
