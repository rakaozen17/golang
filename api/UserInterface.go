package userRouter

import "github.com/labstack/echo/v4"

type User interface {
	RegisterUser(c echo.Context) error
	Login(c echo.Context) error
}
