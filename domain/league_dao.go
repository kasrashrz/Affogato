package domain

import (
	"fmt"
	"github.com/kasrashrz/Affogato/logger"
	"github.com/kasrashrz/Affogato/utils/errors"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func (league *League) Create() *errors.RestErr {
	if err := db.Create(&league).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (league *League) AutoLeagueGenerate() *errors.RestErr {
	var teams []uint
	var teamsCount int
	dao := Team{}
	if err := db.Select("count(id)").Table("teams").Where("league_id IS NULL").Scan(&teamsCount).Error; err != nil {
		teamsCount = 0
	}
	if teamsCount != 0 {
		quantity := 16 - teamsCount
		err := dao.CreateRandom(quantity)
		if err != nil {
			log.Println(err)
		}
	}
	if err := db.Table("teams").Select("id").Where("league_id IS NULL").Limit(16).Find(&teams).Error; err != nil {
		return errors.InternalServerError("something went wrong when deviding 16 teams")
	}

	var match Match
	var matches []Match
	//var finalMatches []Match
	var matchCount = 0

	//t := time.Now()

	if teams != nil {
		if len(teams) == 16 {
			outputLeague := League{
				Model:      gorm.Model{},
				LeagueName: "League",
				LeagueRate: 9,
			}

			if err := db.Create(&outputLeague).Error; err != nil {
				return errors.InternalServerError("something went wrong while creating leagues")
			}

			for _, i2 := range teams {
				if err := db.Table("teams").
					Where("id", i2).
					Updates(map[string]interface{}{"league_id": outputLeague.ID}).Error; err != nil {
					return errors.InternalServerError("something went wrong while updating teams")
				}

				for i4 := 0; i4 < len(teams); i4++ {
					if i2 != teams[i4] {
						match.TeamOneID = i2
						match.TeamTwoID = teams[i4]
						matches = append(matches, match)
					}
				}
			}

			counter := 0

			//for _, m := range matches {
			//	m.MatchTimeDb = t.AddDate(0, 0, counter+1)
			//	finalMatches = append(finalMatches, m)
			//	counter++
			//	if int(counter%(len(teams)+1)) == 0 {
			//		t = t.AddDate(0, 0, 7)
			//		t = t.Add(10 * time.Second)
			//		counter = 0
			//	}
			//	d := 5 * time.Minute
			//	m.MatchTimeDb = m.MatchTimeDb.Round(d)
			//	m.LeagueId = outputLeague.ID
			//
			//	t1 := Team{}
			//	t2 := Team{}
			//	t1.ID = m.TeamOneID
			//	t2.ID = m.TeamTwoID
			//	m.Teams = []Team{t1, t2}
			//	m.MatchType = 1
			//	m.IsAccepted = 0
			//	if err := db.Create(&m).Error; err != nil {
			//		return errors.InternalServerError("something went wrong while matchmaking")
			//	}
			//
			//	//fmt.Println(m.TeamOneID, m.TeamTwoID, m.MatchTimeDb.Round(d))
			//}

			totalMatches := 240
			numDays := 90
			matchesPerDay := totalMatches / numDays

			// Assign match times
			matchDuration := 24 * time.Hour / time.Duration(matchesPerDay)
			currentMatchTime := time.Now()

			for i := range matches {
				dayIndex := counter / matchesPerDay
				matchIndex := counter % matchesPerDay
				matchTime := currentMatchTime.AddDate(0, 0, dayIndex)
				matchTime = matchTime.Add(time.Duration(matchIndex) * matchDuration)
				matches[i].MatchTimeDb = matchTime
				counter++
			}

			for _, m := range matches {
				d := 5 * time.Minute
				m.MatchTimeDb = m.MatchTimeDb.Round(d)
				m.LeagueId = outputLeague.ID
				t1 := Team{}
				t2 := Team{}
				t1.ID = m.TeamOneID
				t2.ID = m.TeamTwoID
				m.Teams = []Team{t1, t2}
				m.MatchType = 1
				m.IsAccepted = 0
				if err := db.Create(&m).Error; err != nil {
					return errors.InternalServerError("something went wrong while matchmaking")
				}
			}

			cup := Cup{
				Model:       gorm.Model{},
				Name:        "hazfi " + strconv.Itoa(rand.Int()),
				MaxCapacity: 16,
				Invites:     16,
				Type:        true,
			}
			db.Create(&cup)

			for i, _ := range teams {
				db.Table("teams").
					Where("id = ?", teams[i]).
					Updates(map[string]interface{}{"cup_h_id": cup.ID})
			}

			for i, id := range teams {
				i += 1
				matchTime := time.Now().AddDate(0, 0, i).String()
				testMatch := Match{
					Model:        gorm.Model{},
					CupId:        cup.ID,
					TeamOneID:    uint(id),
					TeamTwoID:    uint(teams[i+1]),
					TeamOneGoals: 0,
					TeamTwoGoals: 0,
					CupHId:       cup.ID,
					MatchTime:    matchTime,
					IsDone:       0,
					MatchType:    0,
					Priority:     i,
					MatchTimeDb:  time.Now().AddDate(0, 0, 1),
				}
				//fmt.Println("possible matches", id, teamIds[i+1])
				if err := db.Create(&testMatch).Error; err != nil {
					return errors.InternalServerError(err.Error())
				}
				matchCount += 1
				if matchCount == len(teams)/2 {
					return nil
				}
			}
		}
	}

	return nil
}

/*TODO: Increase teams score and coin when draw and win condition*/
func (league *League) MatchResultDecider() *errors.RestErr {
	var matches []Match
	var cupmatches []Match
	var teamOne Team
	var teamTwo Team
	var userOne uint
	var userTwo uint

	if err := db.Table("matches").
		Where("is_done = 0 and match_time_db < now() and cup_id IS NULL").
		Scan(&matches).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if err := db.Table("matches").
		Where("is_done = 0 and match_time_db < now() and (cup_id IS NOT NULL OR cup_h_id IS NOT NULL)").
		Scan(&cupmatches).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if matches == nil && cupmatches == nil {
		return nil
	}

	for _, match := range matches {
		db.Table("teams").Where("id = ?", match.TeamOneID).Scan(&teamOne)
		db.Table("teams").Where("id = ?", match.TeamTwoID).Scan(&teamTwo)
		teamOnePowerAvg, err := teamOne.TeamPowerAvg(int64(teamOne.ID))
		teamOnePowerAvg = (teamOnePowerAvg + 5) * int64(rand.Intn(100))
		teamTwoPowerAvg, err := teamOne.TeamPowerAvg(int64(teamTwo.ID))
		teamTwoPowerAvg = (teamTwoPowerAvg + 5) * int64(rand.Intn(100))

		if err := db.Select("id").Table("users").
			Where("team_id = ?", teamOne.ID).
			Scan(&userOne).Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}
		if err := db.Select("id").Table("users").
			Where("team_id = ?", teamTwo.ID).
			Scan(&userTwo).Error; err != nil {
			return errors.InternalServerError("something went wrong")
		}

		if err != nil {
			return errors.InternalServerError("something went wrong")
		}

		if teamOnePowerAvg == teamTwoPowerAvg {
			rand.Seed(time.Now().UnixNano())
			goals := rand.Intn(5)
			db.Table("matches").
				Where("id = ?", match.ID).
				Updates(map[string]interface{}{"winner_id": nil, "is_done": 1, "team_one_goals": goals, "team_two_goals": goals})
			db.Table("teams").
				Where("id = ?", teamOne.ID).
				Update("score", gorm.Expr("score + ?", 1))
			db.Table("teams").
				Where("id = ?", teamTwo.ID).
				Update("score", gorm.Expr("score + ?", 1))
		}

		if teamOnePowerAvg > teamTwoPowerAvg {
			rand.Seed(time.Now().UnixNano())
			teamOneGoals := rand.Intn(6-3) + 3
			rand.Seed(time.Now().UnixNano())
			teamTwoGoals := rand.Intn(2-0) + 0
			db.Table("matches").
				Where("id = ?", match.ID).
				Updates(map[string]interface{}{"winner_id": teamOne.ID, "is_done": 1, "team_one_goals": teamOneGoals, "team_two_goals": teamTwoGoals})
			db.Table("teams").
				Where("id = ?", teamOne.ID).
				Update("score", gorm.Expr("score + ?", 3))
		}

		if teamOnePowerAvg < teamTwoPowerAvg {
			rand.Seed(time.Now().UnixNano())
			teamTwoGoals := rand.Intn(6-3) + 3
			rand.Seed(time.Now().UnixNano())
			teamOneGoals := rand.Intn(2-0) + 0
			db.Table("matches").
				Where("id = ?", match.ID).
				Updates(map[string]interface{}{"winner_id": teamTwo.ID, "is_done": 1, "team_one_goals": teamOneGoals, "team_two_goals": teamTwoGoals})
			db.Table("teams").
				Where("id = ?", teamTwo.ID).
				Update("score", gorm.Expr("score + ?", 3))
		}

		teamOne = Team{}
		teamTwo = Team{}
		teamOnePowerAvg = 0
		teamTwoPowerAvg = 0
	}

	for _, cupMatch := range cupmatches {
		rand.Seed(time.Now().UnixNano())
		db.Table("teams").Where("id = ?", cupMatch.TeamOneID).Scan(&teamOne)
		db.Table("teams").Where("id = ?", cupMatch.TeamTwoID).Scan(&teamTwo)
		teamOnePowerAvg, err := teamOne.TeamPowerAvg(int64(teamOne.ID))
		teamTwoPowerAvg, err := teamOne.TeamPowerAvg(int64(teamTwo.ID))
		teamOnePowerAvg = (teamOnePowerAvg + 2) * int64(rand.Intn(100))
		rand.Seed(time.Now().UnixNano())
		teamTwoPowerAvg = (teamTwoPowerAvg + 2) * int64(rand.Intn(100))
		fmt.Println(teamOnePowerAvg, teamTwoPowerAvg)
		if err != nil {
			return errors.InternalServerError("something went wrong")
		}

		if teamOnePowerAvg == teamTwoPowerAvg {
			rand.Seed(time.Now().UnixNano())
			goals := rand.Intn(5)
			db.Table("matches").
				Where("id = ?", cupMatch.ID).
				Updates(map[string]interface{}{"winner_id": teamOne.ID, "is_done": 1, "team_one_goals": goals, "team_two_goals": goals})
		}

		if teamOnePowerAvg > teamTwoPowerAvg {
			rand.Seed(time.Now().UnixNano())
			teamOneGoals := rand.Intn(6-3) + 3
			rand.Seed(time.Now().UnixNano())
			teamTwoGoals := rand.Intn(2-0) + 0
			db.Table("matches").
				Where("id = ?", cupMatch.ID).
				Updates(map[string]interface{}{"winner_id": teamOne.ID, "is_done": 1, "team_one_goals": teamOneGoals, "team_two_goals": teamTwoGoals})
		}

		if teamOnePowerAvg < teamTwoPowerAvg {
			rand.Seed(time.Now().UnixNano())
			teamTwoGoals := rand.Intn(6-3) + 3
			rand.Seed(time.Now().UnixNano())
			teamOneGoals := rand.Intn(2-0) + 0
			db.Table("matches").
				Where("id = ?", cupMatch.ID).
				Updates(map[string]interface{}{"winner_id": teamTwo.ID, "is_done": 1, "team_one_goals": teamOneGoals, "team_two_goals": teamTwoGoals})
		}

		teamOne = Team{}
		teamTwo = Team{}
		teamOnePowerAvg = 0
		teamTwoPowerAvg = 0
	}

	return nil
}

func (league *League) CupMatchCreator() *errors.RestErr {
	var cupmatches []Match
	var newMatches []Match
	if err := db.Table("matches").
		Where("is_done = 1 and match_time_db < now() and (cup_id IS NOT NULL OR cup_h_id IS NOT NULL) " +
			"and need_next = 1 and winner_id IS NOT NULL and priority < 15 order by cup_id, priority desc").
		Scan(&cupmatches).Error; err != nil {
		return errors.InternalServerError("something went wrong")
	}

	if cupmatches == nil {
		return nil
	}

	counter := 0
	priorityCounter := 1
	check := cupmatches[0].CupId
	checkCounter := cupmatches[0].Priority

	for i := 0; i <= len(cupmatches)/2; i++ {
		fmt.Println("I ->", i)
		fmt.Println("C ->", counter)
		if counter+1 <= len(cupmatches) {
			if cupmatches[counter].CupId == cupmatches[counter+1].CupId {
				fmt.Println("((", cupmatches[counter].CupId, cupmatches[counter+1].CupId, cupmatches[counter].Priority, "))")

				if check != cupmatches[counter].CupId {
					priorityCounter = 1
					check = cupmatches[counter].CupId
					if checkCounter > 0 {
						checkCounter = cupmatches[counter].Priority
					}
					fmt.Println("SWITCH", cupmatches[counter].Priority)
				}

				match := Match{
					Model:        gorm.Model{},
					CupId:        cupmatches[counter].CupId,
					CupHId:       cupmatches[counter].CupHId,
					TeamOneID:    cupmatches[counter].WinnerID,
					TeamTwoID:    cupmatches[counter+1].WinnerID,
					TeamOneGoals: 0,
					TeamTwoGoals: 0,
					IsDone:       0,
					MatchType:    1,
					Priority:     checkCounter + priorityCounter,
					NeedNext:     1,
					MatchTimeDb:  time.Now().AddDate(0, 0, 1),
				}

				counter += 2
				priorityCounter += 1
				newMatches = append(newMatches, match)
				match = Match{}
			} else {
				fmt.Println("CHANGE")
			}
		}
	}

	for _, match := range newMatches {
		if err := db.Create(&match).Error; err != nil {
			logger.Error(err.Error(), err)
		}
	}

	for _, cupmatch := range cupmatches {
		db.Table("matches").
			Where("id = ?", cupmatch.ID).
			Updates(map[string]interface{}{"need_next": 0})
	}

	return nil
}

func (league *League) LeaderBoard(uid int64) ([]map[string]interface{}, *errors.RestErr) {
	var teams []map[string]interface{}
	var leagueId int64

	if err := db.Select("leagues.id").Table("leagues").
		Joins("join teams t on leagues.id = t.league_id").
		Joins("join users u on t.id = u.team_id").
		Where("u.id = ?", uid).
		Find(&leagueId).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}

	if err := db.Select("id, name, score").Table("teams").Where("league_id = ?", leagueId).
		Order("score desc").Find(&teams).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return teams, nil
}

func (league *League) TodayMatches(userId int64) ([]map[string]interface{}, *errors.RestErr) {

	var todayMatches []map[string]interface{}
	if err := db.Select("m.id, match_time_db, team_one_id, team_two_id").
		Table("users").
		Joins("join teams t on t.id = users.team_id").
		Joins("join leagues l on l.id = t.league_id").
		Joins("join matches m on l.id = m.league_id").
		Where("users.id = ? and datediff(match_time_db, now()) = 0;", userId).
		Find(&todayMatches).
		Error; err != nil {
		return nil, errors.InternalServerError("something went wrong")
	}
	return todayMatches, nil
}
