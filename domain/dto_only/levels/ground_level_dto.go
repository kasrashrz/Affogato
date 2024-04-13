package levels

import "gorm.io/gorm"

type GroundLevels struct {
	gorm.Model
	WeeklyPriceType int64 `json:"weekly_price_type" gorm:"default:1;"`
	WeeklyPrice     int64 `json:"weekly_price" gorm:"default:1;"`
}
