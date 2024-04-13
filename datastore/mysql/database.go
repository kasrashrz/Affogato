package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kasrashrz/Affogato/logger"
)

var (
	Client   *sql.DB
	username = "admin"
	password = "admin123"
	host     = "127.0.0.1:3306"
	db       = "affogato"
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		db,
	)
	logger.Info("about to connect to database")
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		logger.Error("database error", err)
	}
	if err = Client.Ping(); err != nil {
		logger.Error("database error", err)
	}
	logger.Info("connected to database")

}
