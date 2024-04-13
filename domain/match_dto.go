package domain

import (
	"gorm.io/gorm"
	"time"
)

type Match struct {
	gorm.Model
	League       League    `json:"-"`
	LeagueId     uint      `json:"league_id" gorm:"default:null;"`
	Cup          Cup       `json:"-"`
	CupId        uint      `json:"cup_id" gorm:"default:null;"`
	CupH         Cup       `json:"-"`
	CupHId       uint      `json:"cup_h_id" gorm:"default:null;"`
	Teams        []Team    `json:"-" gorm:"many2many:team_matches;"`
	TeamOneID    uint      `json:"team_one_id" gorm:"default:null"`
	TeamOne      Team      `json:"-" gorm:"foreignKey:TeamOneID"`
	TeamTwoID    uint      `json:"team_two_id" gorm:"default:null"`
	TeamTwo      Team      `json:"-" gorm:"foreignKey:TeamTwoID"`
	TeamOneGoals int       `json:"team_one_goals" gorm:"default:0"`
	TeamTwoGoals int       `json:"team_two_goals" gorm:"default:0"`
	WinnerID     uint      `json:"winner_id" gorm:"default:null"`
	Team         Team      `json:"-" gorm:"foreignKey:WinnerID"`
	MatchTime    string    `json:"match_time_json" gorm:"-"`
	IsDone       int       `json:"is_done" gorm:"default:0;"`
	NeedNext     int       `json:"need_next" gorm:"default:1;"`
	MatchType    int       `json:"match_type" gorm:"default:1;"`
	IsAccepted   int       `json:"is_accepted" gorm:"default:0;"`
	IsFriendly   int       `json:"is_friendly" gorm:"default:0;"`
	Priority     int       `json:"priority"`
	MatchTimeDb  time.Time `json:"match_time_db"`
}

type MatchNotification struct {
	ID          uint      `json:"id"`
	TeamOneID   uint      `json:"team_one_id" gorm:"default=NULL"`
	TeamTwoID   uint      `json:"team_two_id" gorm:"default=NULL"`
	MatchTimeDb time.Time `json:"match_time_db"`
}
