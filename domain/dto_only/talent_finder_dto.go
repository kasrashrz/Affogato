package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
)

type TalentFinder struct {
	gorm.Model
	ContractExp         string                   `json:"contract_exp"`
	WeeklySalary        int64                    `json:"weekly_salary"`
	TalentFinderLevel   levels.TalentFinderLevel `json:"talent_finder_level"`
	TalentFinderLevelId uint                     `json:"talent_finder_level_id" gorm:"not null;"`
}
