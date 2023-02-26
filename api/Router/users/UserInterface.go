package userRouter

import "github.com/labstack/echo/v4"

type User interface {
	RegisterUser(c echo.Context) error
	Auth(c echo.Context) error
}
