package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
)

type AssistantCoach struct {
	gorm.Model
	AssistantCoachLevel   levels.AssistantCoachLevel `json:"assistant_coach_level"`
	AssistantCoachLevelId uint                       `json:"assistant_coach_level_id" gorm:"not null;"`
	Practices             int                        `json:"practices"`
}
