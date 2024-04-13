package domain

import (
	"gorm.io/gorm"
)

type Bid struct {
	gorm.Model
	Player   Player `json:"-"`
	PlayerID uint   `json:"pid"`
	Seller   User   `json:"-"`
	SellerID uint   `json:"seller_id"`
	Buyer    User   `json:"-"`
	BuyerID  uint   `json:"buyer_id" gorm:"default:null;"`
	Price    int64  `json:"price"`
	IsDone   bool   `json:"is_done" gorm:"default:0;"`
	EXP      int64  `json:"exp"`
}
