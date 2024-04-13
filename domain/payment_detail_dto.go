package domain

import "gorm.io/gorm"

type PaymentDetail struct {
	gorm.Model
	Description string `json:"description"`
}
