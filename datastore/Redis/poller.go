package Redis

import (
	"errors"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/logger"
	"go.uber.org/zap"
	"gopkg.in/robfig/cron.v2"
	"log"
)

func CronJob() {
	c := cron.New()
	league := domain.League{}
	team := domain.Team{}
	playerPower := domain.PlayerPower{}
	player := domain.Player{}

	//MATCHRESULTDECIDER
	c.AddFunc("@every 5m", func() {
		logger.Info("Match Result Decider started")
		err := league.MatchResultDecider()
		if err != nil {
			log.Println(err)
			logger.Error(err.Message, errors.New(err.Error), zap.Field{
				Key:       "",
				Type:      0,
				Integer:   0,
				String:    "MATCH RESULT DECIDER",
				Interface: nil,
			})
		}
		logger.Info("Match Result Decider ended")
	})

	//BID
	c.AddFunc("@every 7m", func() {
		logger.Info("Bid started")
		err := player.BidCron()
		if err != nil {
			log.Println(err)
			logger.Error(err.Message, errors.New(err.Error), zap.Field{
				Key:       "",
				Type:      0,
				Integer:   0,
				String:    "BID",
				Interface: nil,
			})
		}
		logger.Info("Bid ended")
	})

	//CHECKTRAINERS
	c.AddFunc("@every 3m", func() {
		logger.Info("Check Trainers started")
		err := playerPower.CheckTrainers()
		if err != nil {
			log.Println(err)
			logger.Error(err.Message, errors.New(err.Error), zap.Field{
				Key:       "",
				Type:      0,
				Integer:   0,
				String:    "TRAINER RESET",
				Interface: nil,
			})
		}
		logger.Info("Check Trainers ended")
	})

	//CUPMATCHCREATOR
	c.AddFunc("@every 11m", func() {
		logger.Info("Cup Match Creator started")
		err := league.CupMatchCreator()
		if err != nil {
			log.Println(err)
			logger.Error(err.Message, errors.New(err.Error), zap.Field{
				Key:       "",
				Type:      0,
				Integer:   0,
				String:    "CUP MATCH CREATOR",
				Interface: nil,
			})
		}
		logger.Info("Cup Match Creator ended")
	})

	//PAYMENT
	c.AddFunc("@every 30m", func() {
		logger.Info("Payment started")
		err := team.Payments()
		if err != nil {
			log.Println(err)
			logger.Error(err.Message, errors.New(err.Error), zap.Field{
				Key:       "",
				Type:      0,
				Integer:   0,
				String:    "PAYMENTS",
				Interface: nil,
			})
		}
		logger.Info("Payment ended")
	})

	//TEAMSCOUNT
	//c.AddFunc("@every 2s", func() {
	//	logger.Info("TeamsCount started")
	//	var teamDao domain.Team
	//	teams := teamDao.TeamsCount()
	//	if teams != 0 {
	//		quantity := 16 - teams
	//		err := teamDao.CreateRandom(quantity)
	//		if err != nil {
	//			log.Println(err)
	//			logger.Error(err.Message, errors.New(err.Error), zap.Field{
	//				Key:       "",
	//				Type:      0,
	//				Integer:   0,
	//				String:    "TEAMS COUNT (CREATE RANDOM)",
	//				Interface: nil,
	//			})
	//		}
	//	}
	//	logger.Info("Teams Count ended")
	//})

	//LEAGUEGENERATE
	c.AddFunc("@every 2m", func() {
		logger.Info("Auto League Generate started")
		err := league.AutoLeagueGenerate()
		if err != nil {
			log.Println(err)
			logger.Error(err.Message, errors.New(err.Error), zap.Field{
				Key:       "",
				Type:      0,
				Integer:   0,
				String:    "AUTO LEAGUE GENERATE",
				Interface: nil,
			})
		}
		logger.Info("Auto League Generate ended")
	})

	c.Start()
}
