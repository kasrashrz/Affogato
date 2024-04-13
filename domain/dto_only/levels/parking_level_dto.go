package levels

import "gorm.io/gorm"

type ParkingLevel struct {
	gorm.Model
	DailyIncomePrice     int64 `json:"daily_income_price" gorm:"default:0;"`
	Capacity             int64 `json:"capacity" gorm:"default:100;"`
	WeeklyPriceType      int64 `json:"weekly_price_type" gorm:"default:1;"`
	WeeklyPrice          int64 `json:"weekly_price" gorm:"default:1;"`
	MinimumScoreRequired int64 `json:"minimum_score_required" gorm:"default:0;"`
}
