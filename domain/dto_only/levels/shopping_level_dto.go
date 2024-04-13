package levels

import "gorm.io/gorm"

type ShoppingLevel struct {
	gorm.Model
	DailyIncomePrice int64 `json:"daily_income_price" gorm:"default:0;"`
	WeeklyPrice      int64 `json:"weekly_price" gorm:"default:1;"`
	WeeklyPriceType  int64 `json:"weekly_price_type" gorm:"default:1;"`
}
