package domain

import (
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"gorm.io/gorm"
	"time"
)

type Team struct {
	gorm.Model
	League           League                  `json:"-"`
	LeagueId         uint                    `json:"league_id" gorm:"default:null;"`
	Cup              Cup                     `json:"-"`
	CupId            uint                    `json:"cup_id" gorm:"default:null;"`
	CupH             Cup                     `json:"-"`
	CupHId           uint                    `json:"cup_h_id" gorm:"default:null;"`
	Stadium          dto_only.Stadium        `json:"stadium" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name             string                  `json:"name"`
	Score            int64                   `json:"score"`
	Sponsor          string                  `json:"sponsor"`
	ChangedName      bool                    `json:"changed_name" gorm:"default: 0"`
	Trainer          dto_only.Trainer        `json:"trainer"`
	TrainerId        uint                    `json:"trainer_id"`
	AssistantCoach   dto_only.AssistantCoach `json:"assistant_coach"`
	AssistantCoachId uint                    `json:"assistant_coach_id"`
	FitnessCoach     dto_only.FitnessCoach   `json:"fitness_coach"`
	FitnessCoachId   uint                    `json:"fitness_coach_id"`
	Doctor           dto_only.Doctor         `json:"doctor"`
	DoctorId         uint                    `json:"doctor_id"`
	TalentFinder     dto_only.TalentFinder   `json:"talent_finder"`
	TalentFinderId   uint                    `json:"talent_finder_id"`
	Tactic           dto_only.Tactic         `json:"tactic" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TacticId         uint                    `json:"tactic_id" gorm:"default:1"`
	LastPaid         time.Time               `json:"last_paid" gorm:"default: NULL"`
	LastEarned       time.Time               `json:"last_earned" gorm:"default: NULL"`
	TacticFormation  string                  `json:"tactic_formation"`
	Players          []Player                `json:"players" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CaptainID        uint                    `json:"captain_id" gorm:"default:null"`
	Player           Player                  `gorm:"foreignKey:CaptainID"`
}

type GetAllTeams struct {
	gorm.Model
	Name  string `json:"name"`
	Score int64  `json:"score"`
	//Sponsor string `json:"sponsor"`
}
