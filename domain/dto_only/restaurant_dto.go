package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
	"time"
)

type Restaurant struct {
	gorm.Model
	RestaurantLevel   levels.RestaurantLevel `json:"restaurant_level"`
	RestaurantLevelId uint                   `json:"restaurant_level_id"`
	StadiumID         uint                   `json:"stadium_id" gorm:"default:null;"`
	LastPaid          time.Time              `json:"last_paid"`
	LastEarned        time.Time              `json:"last_earned"`
}
