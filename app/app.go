package app

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	ctx context.Context
	dbw *sqlx.DB
	dbr *sqlx.DB
}

func InitApp() *App {
	c := &App{
		ctx: context.Background(),
	}
	logrus.Info("Initialize Config for the application")
	c.InitConfig()

	return c
}

func (app *App) InitConfig() {
	viper.SetConfigName("config")

	// Set the configuration file type
	viper.SetConfigType("yaml")

	// Set the configuration file path
	viper.AddConfigPath(".")

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

func (app *App) CloseApp() {
	logrus.Info("Closing the database connection..")
	if app.dbw != nil {
		app.dbw.Close()
	}
	if app.dbr != nil {
		app.dbr.Close()
	}

	logrus.Info("Ending worker...")
}
