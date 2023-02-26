package userRouter

import (
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, s User, m ...echo.MiddlewareFunc) {
	users := e.Group("v1/users")
	users.POST("/register", s.RegisterUser)
	users.POST("/login", s.Auth)
}
