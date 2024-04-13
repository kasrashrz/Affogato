package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
	"time"
)

type Transportation struct {
	gorm.Model
	TransportationLevel   levels.TransportationLevel `json:"transportation_level"`
	TransportationLevelId uint                       `json:"transportation_level_id" default:"null"`
	StadiumID             uint                       `json:"stadium_id" gorm:"default:null;"`
	LastPaid              time.Time                  `json:"last_paid"`
	LastEarned            time.Time                  `json:"last_earned"`
}
