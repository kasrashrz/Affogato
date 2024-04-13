package dto_only

import "gorm.io/gorm"

type Tactic struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null;"`
	TacticFormat string `json:"tactic" gorm:"not null;"`
	Strategy     string `json:"strategy" gorm:"default:null;"`
}

type GetAllTactics struct {
	ID           uint   `json:"id"`
	Name         string `json:"name" gorm:"not null;"`
	TacticFormat string `json:"tactic" gorm:"not null;"`
	Strategy     string `json:"strategy" gorm:"default:null;"`
}
