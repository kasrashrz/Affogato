package domain

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	PostName string `json:"post_name"`
}
