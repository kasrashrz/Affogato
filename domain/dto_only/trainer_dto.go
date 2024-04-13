package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
)

type Trainer struct {
	gorm.Model
	ContractExp                     string              `json:"contract_exp"`
	WeeklySalary                    int64               `json:"weekly_salary"`
	TrainerLevel                    levels.TrainerLevel `json:"trainer_level"`
	TrainerLevelId                  uint                `json:"trainer_level_id" gorm:"not null;"`
	RemainingDailyPracticeDuration  int                 `json:"remaining_daily_practice_duration" gorm:"default:5"`
	RemainingPlayerPracticeDuration int                 `json:"remaining_player_practice_duration" gorm:"default:5"`
	RemainingExtraPracticeDuration  int                 `json:"remaining_extra_practice_duration" gorm:"default:5"`
}
