package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
)

type FitnessCoach struct {
	gorm.Model
	ContractExp         string                   `json:"contract_exp"`
	WeeklySalary        int64                    `json:"weekly_salary"`
	FitnessCoachLevel   levels.FitnessCoachLevel `json:"fitness_coach_level"`
	FitnessCoachLevelId uint                     `json:"fitness_coach_level_id" gorm:"not null;"`
}
