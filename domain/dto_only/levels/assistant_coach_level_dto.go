package levels

import "gorm.io/gorm"

type AssistantCoachLevel struct {
	gorm.Model
	Salary                 int64 `json:"salary"`
	DailyPracticeDuration  int   `json:"daily_practice_duration" gorm:"default:2"`
	PlayerPracticeDuration int   `json:"player_practice_duration" gorm:"default:2"`
	ExtraPracticeDuration  int   `json:"extra_practice_duration" gorm:"default:2"`
}
