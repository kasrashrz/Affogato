package dto_only

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	StatusName string `json:"status_name"`
}
