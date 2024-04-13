package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
)

type Lights struct {
	gorm.Model
	LightsLevel   levels.LightLevels `json:"lights_level"`
	LightsLevelId uint               `json:"light_level_id" gorm:"default:null"`
	StadiumId     uint               `json:"stadium_id" gorm:"default:null"`
	Price         int64              `json:"price"`
	RequiredScore int64              `json:"required_score"`
	IncreaseFans  int64              `json:"increase_fans"`
}
