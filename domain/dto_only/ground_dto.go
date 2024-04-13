package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
)

type Ground struct {
	gorm.Model
	GroundLevel   levels.GroundLevels `json:"ground_level"`
	GroundLevelID int                 `json:"ground_level_id"`
	RequiredScore int64               `json:"required_score"`
	StadiumID     uint                `json:"stadium_id" gorm:"default:null;"`
}
