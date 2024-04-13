package levels

import "gorm.io/gorm"

type TalentFinderLevel struct {
	gorm.Model
	ObservePlayers         int   `json:"observe_players"`
	ObservationTime        int   `json:"observation_time"`
	EncryptionTransfer     bool  `json:"encryption_transfer" gorm:"default:0"`
	WeeklyTransferCapacity int   `json:"weekly_transfer_capacity"`
	Salary                 int64 `json:"salary"`
	MaxForeignerPlayers    int   `json:"max_foreigner_players"`
}
