package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	DoctorLevel   levels.DoctorLevel `json:"doctor_level"`
	DoctorLevelId uint               `json:"doctor_level_id" gorm:"not null;"`
	ContractExp   string             `json:"contract_exp"`
	WeeklySalary  int64              `json:"weekly_salary"`
}
