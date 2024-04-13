package domain

import (
	"testing"
)

func TestCreateRandom(t *testing.T) {
	var teamsCount int
	if err := db.Select("count(id)").Table("teams").Scan(&teamsCount).Error; err != nil {
		t.Errorf("database error")
	}
}
