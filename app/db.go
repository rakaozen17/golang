package app

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (app *App) DbWrite() *sqlx.DB {
	Driver := viper.GetString("database.driver")
	Host := viper.GetString("database.host")
	Port := viper.GetInt("database.port")
	Username := viper.GetString("database.username")
	Password := viper.GetString("database.password")
	Database := viper.GetString("database.database")
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", Host, Username, Password, Port, Database)
	conn, err := sqlx.Open(Driver, connString)
	if err != nil {
		logrus.Panic(err)
	} else {
		logrus.Infof("Successfully connected to database %v", Database)
	}

	return conn
}

func (app *App) DbRead() *sqlx.DB {
	return app.DbWrite()
}
