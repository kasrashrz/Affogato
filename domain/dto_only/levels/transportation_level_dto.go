package levels

import "gorm.io/gorm"

type TransportationLevel struct {
	gorm.Model
	WeeklyPrice     int64 `json:"weekly_price" gorm:"default:1;"`
	WeeklyPriceType int64 `json:"weekly_price_Type" gorm:"default:1;"`
}
