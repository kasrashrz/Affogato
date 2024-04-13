package domain

import (
	"fmt"
	_ "fmt"
	"github.com/kasrashrz/Affogato/configs"
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"github.com/kasrashrz/Affogato/domain/dto_only/levels"
	"github.com/kasrashrz/Affogato/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func SetupModels() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.ReadConfig().Database.User,
		configs.ReadConfig().Database.Password,
		configs.ReadConfig().Database.Host,
		configs.ReadConfig().Database.Port,
		configs.ReadConfig().Database.Name,
	)
	localDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: false})
	if err != nil {
		logger.Error("gorm problem", err)
	}
	if err := localDb.AutoMigrate(&User{}, &Team{}, &Player{},
		&dto_only.Ground{}, &dto_only.Stadium{}, &PlayerPower{},
		&Match{}, &Post{}, &dto_only.Status{}, &League{},
		&dto_only.Trainer{}, &levels.TrainerLevel{},
		&dto_only.Tactic{}, &dto_only.AssistantCoach{},
		&dto_only.Lights{}, &levels.LightLevels{}, &levels.AssistantCoachLevel{},
		&dto_only.FitnessCoach{}, &levels.FitnessCoachLevel{},
		&dto_only.Doctor{}, &levels.DoctorLevel{}, &levels.ScoreBoardLevels{},
		&dto_only.ScoreBoard{}, &levels.ParkingLevel{}, &levels.RestaurantLevel{},
		&dto_only.Parking{}, &dto_only.Restaurant{}, &PaymentDetail{},
		&PaymentHistory{}, &Market{}, &levels.ShoppingLevel{},
		&dto_only.Shopping{}, &levels.TransportationLevel{},
		&dto_only.Transportation{}, &Mission{}, levels.StadiumLevel{},
		&MissionsTracker{}, &Bid{}, &TransferData{}, &Cup{}); err != nil {
		logger.Error(err.Error(), err)
	}
	db = localDb
	return localDb
}
