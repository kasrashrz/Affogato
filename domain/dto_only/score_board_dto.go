package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
)

type ScoreBoard struct {
	gorm.Model
	ScoreBoardLevel   levels.ScoreBoardLevels `json:"score_board_level"`
	ScoreBoardLevelID uint                    `json:"score_board_level_id"`
	Price             int64                   `json:"price"`
	RequiredScore     int64                   `json:"required_score"`
	StadiumID         uint                    `json:"stadium_id" gorm:"default:null;"`
}
