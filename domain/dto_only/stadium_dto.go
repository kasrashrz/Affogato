package dto_only

import (
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"gorm.io/gorm"
)

type Stadium struct {
	gorm.Model
	TotalCapacity  int64               `json:"total_capacity"`
	TicketPrice    int64               `json:"ticket_price"`
	StadiumLevel   levels.StadiumLevel `json:"-"`
	StadiumLevelId uint                `json:"level" gorm:"default:1;"`
	Pleasure       int                 `json:"pleasure"`
	Ground         Ground              `json:"ground" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ScoreBoard     ScoreBoard          `json:"score-board" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Lights         Lights              `json:"lights" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Parking        Parking             `json:"parking" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Restaurant     Restaurant          `json:"restaurant" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Shopping       Shopping            `json:"shopping" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Transportation Transportation      `json:"transportation" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TeamID         uint                `json:"team_id" gorm:"default:null;"`
	FansAmount     int64               `json:"fans_amount"`
}
