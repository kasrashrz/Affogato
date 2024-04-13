package player_utils

import (
	"github.com/kasrashrz/Affogato/domain"
)

func PowerCheck(power domain.PlayerPower) bool {
	if power.Power < 25 || power.Control < 25 || power.Dribble < 25 ||
		power.Endurance < 25 || power.Goal < 25 || power.Head < 25 || power.Pass < 25 || power.Shoot < 25 {
		return false
	}
	return true
}
