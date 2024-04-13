package levels

import "gorm.io/gorm"

type StadiumLevel struct {
	gorm.Model
	Price                 int64  `json:"price"`
	PriceType             string `json:"price_type" gorm:"default:C"`
	FirstCapacity         int64  `json:"first_capacity"`
	IncreaseSeatsAmount   int64  `json:"increase_seats_amount"`
	PricePerThousandSeats int64  `json:"price_per_each_seat"`
}
