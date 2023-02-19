package userRouter

import (
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, s User, m ...echo.MiddlewareFunc) {
	crypto := e.Group("v1/users")
	crypto.POST("/register", s.RegisterUser)
	crypto.POST("/login", s.Login)
}
