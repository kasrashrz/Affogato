package domain

import "gorm.io/gorm"

type Mission struct {
	gorm.Model
	Users []User `gorm:"many2many:users_missions;"`
	Name  string `json:"name"`
	Prize int64  `json:"prize"`
}
