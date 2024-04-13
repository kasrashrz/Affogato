package levels

import "gorm.io/gorm"

type TrainerLevel struct {
	gorm.Model
	Salary                 int64 `json:"salary"`
	Power                  int   `json:"power" gorm:"default:5"`
	DailyPracticeDuration  int   `json:"daily_practice_duration" gorm:"default:5"`
	PlayerPracticeDuration int   `json:"player_practice_duration" gorm:"default:5"`
	ExtraPracticeDuration  int   `json:"extra_practice_duration" gorm:"default:5"`
}
