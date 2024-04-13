package levels

import "gorm.io/gorm"

type FitnessCoachLevel struct {
	gorm.Model
	Salary              int64 `json:"salary"`
	PowerIncreasePerDay int   `json:"power_increase_per_day"`
}
