package domain

import "gorm.io/gorm"

type PlayerPower struct {
	gorm.Model
	PlayerID  uint `json:"player_id" gorm:"default:null;"`
	Control   int  `json:"control" gorm:"default:25;"`
	Power     int  `json:"power" gorm:"default:25;"`
	Pass      int  `json:"pass" gorm:"default:25;"`
	Shoot     int  `json:"shoot" gorm:"default:25;"`
	Dribble   int  `json:"dribble" gorm:"default:25;"`
	Tackle    int  `json:"tackle" gorm:"default:25;"`
	Head      int  `json:"head" gorm:"default:25;"`
	Endurance int  `json:"endurance" gorm:"default:25;"`
	Strength  int  `json:"strength" gorm:"default:25;"`
	Goal      int  `json:"goal" gorm:"default:25;"`
}
