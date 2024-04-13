package domain

import "gorm.io/gorm"

type MissionsTracker struct {
	gorm.Model
	User      User
	Mission   Mission
	UserId    uint  `json:"user_id" gorm:"default:null;"`
	MissionId uint  `json:"mission_id" gorm:"default:null;"`
	Count     int64 `json:"count"`
}
