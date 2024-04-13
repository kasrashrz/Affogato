package main

import (
	"github.com/kasrashrz/Affogato/app"
	"github.com/kasrashrz/Affogato/datastore/Redis"
	"github.com/kasrashrz/Affogato/domain"
	"github.com/kasrashrz/Affogato/logger"
)

func main() {
	logger.Info("Server is about to start")
	Redis.CronJob()
	domain.SetupModels()
	app.StartServer()
}

//TODO: player search from market -> DONE
//TODO: player search from list -> DONE
//TODO: league generate -> DONE
//TODO: league game generate -> DONE
//TODO: Setup API (FULLY) -> ٍRestaurant, parking, stadium, every levels, etc... -> DONE
//TODO:	Player Practice -> DOONE
//TODO: league games must be done correctly
//TODO:	CRON jobs -> TeamsCount -> DONE
//TODO:	CRON jobs -> Auto League Generate -> DONE
//TODO:	CRON jobs -> Match Result Decider -> DONE
//TODO:	CRON jobs -> Cup Match Creator -> DONE

//TODO: score must be on a team and it must increase
//TODO:	CRON jobs -> Rent Payment
//TODO: more accurate logs
//TODO:	CRON TIMER -> ENV (COMPOSE)
//TODO:	Transfer
//TODO:	Transfer Data
