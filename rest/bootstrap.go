package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	destination "restproject/api/impl/destination"
	users "restproject/api/impl/users"
	"restproject/app"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Serve(application *app.App) {
	// Echo instance
	e := echo.New()
	e.Debug = false
	e.HideBanner = true
	e.HidePort = true
	e.Pre(
		middleware.RemoveTrailingSlash(),
	)

	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.RequestID(),
	)

	//init database yang akan digunakan disini
	SQLTRAVEL_DBR := application.DbRead()
	SQLTRAVEL_DBW := application.DbWrite()

	//REGISTER ROUTER HERE
	users.RegisterRoute(application, e, SQLTRAVEL_DBR, SQLTRAVEL_DBW)
	destination.RegisterRoute(application, e, SQLTRAVEL_DBR, SQLTRAVEL_DBW)

	go func() {
		port := viper.GetString("app.port")
		if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
			if err == http.ErrServerClosed {
				logrus.Info("server stopped")
			} else {
				logrus.Fatal(fmt.Printf("server failed to start on port %s", port))
			}
		} else {
			logrus.Infof("Server started at port %s", port)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	logrus.Info("Shutting down server")
	application.CloseApp()

	if err := e.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		logrus.Fatal(fmt.Printf("failed to gracefully shut down the server with error : %v", err))
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
