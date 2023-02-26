package hotelsRouter

import "github.com/labstack/echo/v4"

type Hotels interface {
	CreateHotel(c echo.Context) error
	UpdateHotel(c echo.Context) error
	DeleteHotel(c echo.Context) error
	GetSingleHotel(c echo.Context) error
	GetAllHotel(c echo.Context) error
}
