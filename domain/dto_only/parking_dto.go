package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
	"time"
)

type Parking struct {
	gorm.Model
	ParkingLevel   levels.ParkingLevel `json:"parking_level"`
	ParkingLevelId uint                `json:"parking_level_id" gorm:"default:null;"`
	StadiumID      uint                `json:"stadium_id" gorm:"default:null;"`
	LastPaid       time.Time           `json:"last_paid"`
	LastEarned     time.Time           `json:"last_earned"`
}
