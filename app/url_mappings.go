package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Affogato/controllers/cups_controller"
	"github.com/kasrashrz/Affogato/controllers/jwt"
	"github.com/kasrashrz/Affogato/controllers/leagues_controller"
	"github.com/kasrashrz/Affogato/controllers/markets_controller"
	"github.com/kasrashrz/Affogato/controllers/matches_controller"
	"github.com/kasrashrz/Affogato/controllers/ping"
	"github.com/kasrashrz/Affogato/controllers/players_controller"
	"github.com/kasrashrz/Affogato/controllers/players_power_controller"
	"github.com/kasrashrz/Affogato/controllers/teams_controller"
	"github.com/kasrashrz/Affogato/controllers/users_controller"
	"github.com/kasrashrz/Affogato/domain"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/server-time", ping.ServerTime)
	router.PUT("/setup", ping.DatabaseSetup)
	/**********************************		MISSIONS		*************************************/
	router.GET("/missions", users_controller.GetAllMissions)
	router.GET("/missions/user", users_controller.GetUserMissions)
	/**********************************		USER		*************************************/
	router.GET("/users/token-to-id", users_controller.GetUidFromEmail)
	router.GET("/igm", jwt.HandleGoogleLogin)
	router.POST("/login", jwt.FinalLogin)
	router.GET("/users", users_controller.GetAll)
	router.GET("/users/get-one", users_controller.GetOne)
	router.POST("/users/add", users_controller.Create)
	router.PUT("/users/coin/increase", users_controller.IncreaseCoin)
	router.PUT("/users/coin/decrease", users_controller.DecreaseCoin)
	router.PUT("/users/gem/increase", users_controller.IncreaseGem)
	router.PUT("/users/change/username", users_controller.UpdateUsername)
	router.PUT("/users/gem/decrease", users_controller.DecreaseGem)
	router.PUT("/users/friend/invite", users_controller.AddFriend)
	router.GET("/users/friend/get-all", users_controller.GetFriends)
	router.DELETE("/users/delete", users_controller.Delete)
	router.PUT("/users/avatar/change-formation", users_controller.ChangeAvatarFormat)
	router.GET("/users/avatar/get-formation", users_controller.GetAvatarFormation)
	router.GET("/users/get-data", cups_controller.UserLeagueAndCupData)
	/**********************************		PLAYER		*************************************/
	router.POST("/bids/add", players_controller.AddToBid)
	router.GET("/bids", players_controller.GetAllBids)
	router.GET("/bids/get-one", players_controller.GetOneBid)
	router.PUT("/bids/increase", players_controller.IncreaseBid)
	/**********************************		BID		*************************************/
	router.GET("/players", players_controller.GetAll)
	router.GET("/players/get-one", players_controller.GetOne)
	router.GET("/players/search-with-filter", players_controller.SearchWithFilter)
	router.GET("/players/p2p/search-with-filter", players_controller.SearchP2PWithFilter)
	router.POST("/players/add", players_controller.Create)
	router.GET("/players/without-team", players_controller.GetPlayersWithoutTeam)
	router.GET("/players/random-generate", players_controller.GenerateRandomPlayers)
	router.PUT("/players/increase/goals", players_controller.GoalIncrement)
	router.PUT("/players/set-default/goals", players_controller.SetDefaultGoals)
	router.PUT("/players/age/increase", players_controller.IncreaseAge)
	router.PUT("/players/age/remove-olds", players_controller.RemoveOlds)
	router.PUT("/players/status/change", players_controller.ChangeStatus)
	router.GET("/players/power/avg", players_controller.PlayerAvg)
	router.GET("/players/power/list", players_controller.PowerList)
	router.PUT("/players/body-builder", players_controller.BodyBuilder)
	router.PUT("/players/remove-from-team", players_controller.RemovePlayerFromTeam)
	router.DELETE("/players/delete", players_controller.Delete)
	/*******************		PLAYER POWER		**********************/
	router.PUT("/players/practice", players_power_controller.PracticePlayer)
	router.PUT("/players/power/set-default", players_power_controller.SetDefaultPowers)
	/**********************************		TEAM		*************************************/
	router.POST("/teams/create", teams_controller.Create)
	router.GET("/teams", teams_controller.GetAll)
	router.GET("/teams/get-one", teams_controller.GetOne)
	router.GET("/teams/players", teams_controller.GetPlayers)
	router.GET("/teams/avg/power", teams_controller.TeamPowerAvg)
	router.GET("/teams/avg/energy", teams_controller.TeamEnergyAvg)
	router.DELETE("/teams/delete", teams_controller.Delete)
	router.PUT("/teams/change/score", teams_controller.ChangeScore)
	router.PUT("/teams/change/name", teams_controller.ChangeName)
	router.PUT("/teams/change/captain", teams_controller.ChangeTeamCaptain)
	router.PUT("/teams/change/ticket-price", teams_controller.ChangeTicketPrice)
	router.PUT("/teams/random-generate", teams_controller.RandomGenerate)
	router.GET("/teams/tactics", teams_controller.GetAllTactics)
	router.PUT("/teams/tactic/change-tactic", teams_controller.ChangeTactic)
	router.PUT("/teams/tactic/change-formation", teams_controller.ChangeTacticFormation)
	router.GET("/teams/tactic/opponent", teams_controller.TeamTactic)
	router.GET("/teams/tactic/formation", teams_controller.GetTacticFormation)
	router.GET("/teams/workers", teams_controller.Workers)
	router.GET("/teams/purchase-history", teams_controller.GetPurchaseHistory)
	router.GET("/teams/stadium/status", teams_controller.GetStadiumData)
	router.GET("/teams/stadium/levels", teams_controller.GetStadiumLevels)
	router.GET("/teams/income", teams_controller.ListOfIncome)
	router.GET("/teams/outcome", teams_controller.ListOfOutcome)
	router.GET("/teams/id-to-name", teams_controller.IdToName)
	router.GET("/teams/transfer-data", teams_controller.GetTransferData)
	router.PUT("/teams/remove/doctor", teams_controller.RemoveDoctor)
	router.PUT("/teams/remove/assistant-coach", teams_controller.RemoveAssistantCoach)
	router.PUT("/teams/remove/fitness-coach", teams_controller.RemoveFitnessCoach)
	router.PUT("/teams/remove/talent-finder", teams_controller.RemoveTalentFinder)
	router.PUT("/teams/remove/trainer", teams_controller.RemoveTrainer)
	router.PUT("/teams/upgrade/section", teams_controller.UpgradeSection)
	router.PUT("/teams/upgrade/stadium", teams_controller.UpgradeStadium)
	/**********************************		MATCH		*************************************/
	router.POST("/matches/create", matches_controller.Create)
	router.GET("/matches/notification", matches_controller.MatchNotification)
	router.GET("/matches/users-matches", matches_controller.UsersMatches)
	router.GET("/matches/single-match", matches_controller.SingleMatch)
	router.PUT("/matches/update", matches_controller.UpdateResult)
	/**********************************		FRIENDLY MATCHES		*************************************/
	router.POST("/matches/friendly/create", matches_controller.FriendlyMatch)
	router.GET("/matches/friendly/matches", matches_controller.GetFriendlyMatchesInvites)
	router.PUT("/matches/friendly/join", matches_controller.JoinFriendlyMatch)

	/**********************************		MARKET		*************************************/
	router.GET("/market/available/doctors", teams_controller.GetAllDoctors)
	router.GET("/market/available/trainers", teams_controller.GetAllTrainers)
	router.GET("/market/available/assistant-coaches", teams_controller.GetAllAssistantCoaches)
	router.GET("/market/available/fitness-coaches", teams_controller.GetAllFitnessCoaches)
	router.GET("/market/available/talent-finders", teams_controller.GetAllTalentFinders)
	router.GET("/market/available/players", markets_controller.AllMarketChoices)
	dao := domain.League{}
	router.GET("/test", func(context *gin.Context) {
		dao.CupMatchCreator()
	})
	router.PUT("/market/buy/assistant-coach", teams_controller.BuyAssistantCoach)
	router.PUT("/market/buy/doctor", teams_controller.BuyDoctor)
	router.PUT("/market/buy/fitness-coach", teams_controller.BuyFitnessCoach)
	router.PUT("/market/buy/talent-finder", teams_controller.BuyTalentFinder)
	router.PUT("/market/buy/player", markets_controller.BuyPlayer)
	router.PUT("/market/buy/trainer", teams_controller.BuyTrainer)
	router.POST("/market/sell/player", markets_controller.AddPlayer)
	/**********************************		LEAGUE		*************************************/
	router.POST("/leagues/create", leagues_controller.Create)
	router.POST("/leagues/join", leagues_controller.Create)
	router.GET("/leagues/leader-board", leagues_controller.LeaderBoard)
	router.GET("/leagues/today-matches", leagues_controller.TodayMatches)
	/**********************************		CUP		*************************************/
	router.POST("/cups/create", cups_controller.Create)
	router.GET("/cups", cups_controller.GetAll)
	router.GET("/cups/get-one", cups_controller.GetOne)
	router.GET("/cups/cup-data", cups_controller.GetCupMatches)
	router.PUT("/cups/join", cups_controller.Join)

	/**********************************		TEST API		*************************************/

}
