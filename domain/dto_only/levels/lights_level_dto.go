package levels

import "gorm.io/gorm"

type LightLevels struct {
	gorm.Model
	Max int64 `json:"max"`
}
