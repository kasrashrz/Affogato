package domain

import (
	"fmt"
	"github.com/kasrashrz/Affogato/utils/errors"
	"gorm.io/gorm"
	"time"
)

func (match Match) Create() *errors.RestErr {

	teamOne := Team{
		Model: gorm.Model{
			ID: match.TeamOneID,
		},
	}
	teamTwo := Team{
		Model: gorm.Model{
			ID: match.TeamTwoID,
		},
	}
	newMatch := Match{
		Teams:        []Team{teamOne, teamTwo},
		TeamOneID:    match.TeamOneID,
		TeamTwoID:    match.TeamTwoID,
		TeamOneGoals: match.TeamOneGoals,
		TeamTwoGoals: match.TeamTwoGoals,
		MatchTime:    match.MatchTime,
	}
	newMatch.MatchTimeDb, _ = time.Parse("2006-01-02 15:04:05", match.MatchTime)
	fmt.Println(match.MatchTimeDb.Format("2006-01-02 15:04:05"))
	newMatch.MatchTimeDb.Add(time.Minute * 270)
	if err := db.Create(&newMatch).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (match Match) MatchNotification(uid int64) ([]MatchNotification, *errors.RestErr) {
	var teamId int64
	if err := db.Select("team_id").Table("users").
		Where("id = ?", uid).Scan(&teamId).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	var matches []MatchNotification
	if err := db.Table("matches").
		Joins("join team_matches tm on matches.id = tm.match_id").
		//and match_time_db > now()
		Where("team_id = ?", teamId).
		Order("timestampdiff(SECOND, matches.match_time_db, NOW()) desc").
		Scan(&matches).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return matches, nil
}

func (match Match) Update() *errors.RestErr {

	var winnerId uint
	winnerId = 0
	var findMatch Match
	if err := db.Table("matches").Where("id = ?", match.ID).Scan(&findMatch).Error; err != nil {
		return errors.NotFoundError("match not found")
	}

	if match.TeamTwoGoals > match.TeamOneGoals {
		winnerId = findMatch.TeamTwoID
	}
	if match.TeamTwoGoals < match.TeamOneGoals {
		winnerId = findMatch.TeamOneID
	}

	//queries for handling winnerID
	withWinnerId := db.Table("matches").
		Where("id = ?", match.ID).
		Updates(map[string]interface{}{"team_one_goals": match.TeamOneGoals,
			"team_two_goals": match.TeamTwoGoals,
			"winner_id":      winnerId}).Error
	withoutWinnerId := db.Table("matches").Where("id = ?", match.ID).
		Updates(map[string]interface{}{"team_one_goals": match.TeamOneGoals,
			"team_two_goals": match.TeamTwoGoals}).Error

	if match.TeamTwoGoals == match.TeamOneGoals {
		if withoutWinnerId != nil {
			return errors.InternalServerError("something went wrong")
		}
		//db.Table("matches").Where("id = ?", match.ID).Scan(&history)
		//db.Create(&history)
		db.Exec("delete from team_matches where match_id = ?;", match.ID)
		db.Unscoped().Delete(&match)
		return nil
	}

	if withWinnerId != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (match Match) UsersMatches(id int64) ([]Match, *errors.RestErr) {

	var matches []Match

	if err := db.Table("matches").
		Where("is_done = 1 AND (team_one_id = ? OR team_two_id = ?) ", id, id).Limit(10).
		Order("match_time_db DESC").Find(&matches).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return matches, nil
}

func (match Match) SingleMatch(id int64) (*Match, *errors.RestErr) {

	if err := db.Table("matches").
		Where("id = ?", id).
		Find(&match).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")

	}
	return &match, nil
}

func (testMatch Match) FriendlyMatch(match Match) *errors.RestErr {
	if err := db.Select("team_id").Table("users").Where("id = ?", match.TeamOneID).
		Scan(&match.TeamOneID).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	if err := db.Select("team_id").Table("users").Where("id = ?", match.TeamTwoID).
		Scan(&match.TeamTwoID).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	teamOne := Team{
		Model: gorm.Model{
			ID: match.TeamOneID,
		},
	}
	teamTwo := Team{
		Model: gorm.Model{
			ID: match.TeamTwoID,
		},
	}
	newMatch := Match{
		Teams:        []Team{teamOne, teamTwo},
		TeamOneID:    match.TeamOneID,
		TeamTwoID:    match.TeamTwoID,
		TeamOneGoals: match.TeamOneGoals,
		TeamTwoGoals: match.TeamTwoGoals,
		IsAccepted:   0,
		IsFriendly:   1,
		MatchTime:    match.MatchTime,
	}
	fmt.Println(match.MatchTime)
	var timeErr error
	newMatch.MatchTimeDb, timeErr = time.Parse("2006-01-02 15:04:05", match.MatchTime)
	if timeErr != nil {
		fmt.Println(timeErr)
	}
	fmt.Println(match.MatchTimeDb.Format("2006-01-02 15:04:05"))
	newMatch.MatchTimeDb.Add(time.Minute * 270)
	if err := db.Create(&newMatch).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (match Match) JoinFriendlyMatch(matchId int64) *errors.RestErr {

	if err := db.Table("matches").Where("id = ?", matchId).
		Updates(map[string]interface{}{"is_accepted": 1}).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (match Match) GetFriendlyMatchesInvites(uid int64) ([]MatchNotification, *errors.RestErr) {
	var teamId int64
	if err := db.Select("team_id").Table("users").
		Where("id = ?", uid).Scan(&teamId).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	var matches []MatchNotification
	if err := db.Table("matches").
		Where("is_accepted = 0 AND is_friendly = 1 AND team_two_id = ?", teamId).
		Order("timestampdiff(SECOND, matches.match_time_db, NOW()) desc").
		Scan(&matches).Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	return matches, nil
}

//small change
