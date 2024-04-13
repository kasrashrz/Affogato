package levels

import "gorm.io/gorm"

type DoctorLevel struct {
	gorm.Model
	Healing        int   `json:"healing"`
	MaxTeamPlayers int   `json:"max_team_players"`
	Salary         int64 `json:"salary"`
}
