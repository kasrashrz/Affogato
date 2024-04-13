package domain

import (
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"gorm.io/gorm"
)

type PaymentHistory struct {
	gorm.Model
	Team             Team                    `json:"-"`
	TeamId           int64                   `json:"team_id"`
	PaymentDetail    PaymentDetail           `json:"payment_detail"`
	PaymentDetailId  int64                   `json:"payment_detail_id"`
	AssistantCoach   dto_only.AssistantCoach `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AssistantCoachId int64                   `json:"assistant_coach_id" gorm:"default:null"`
	Doctor           dto_only.Doctor         `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DoctorId         int64                   `json:"doctor_id" gorm:"default:null"`
	FitnessCoach     dto_only.FitnessCoach   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FitnessCoachId   int64                   `json:"fitness_coach_id" gorm:"default:null"`
	TalentFinder     dto_only.TalentFinder   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TalentFinderId   int64                   `json:"talent_finder_id" gorm:"default:null"`
	Trainer          dto_only.Trainer        `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TrainerId        int64                   `json:"trainer_id" gorm:"default:null"`
	Player           Player                  `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PlayerId         int64                   `json:"player_id" gorm:"default:null"`
	Price            int64                   `json:"price"`
	IsBuy            bool                    `json:"buy_or_sell" gorm:"default:1;"`
}
