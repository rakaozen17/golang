package destinationRouter

import "github.com/labstack/echo/v4"

type Destination interface {
	CreateDestination(c echo.Context) error
	UpdateDestination(c echo.Context) error
	DeleteDestination(c echo.Context) error
	GetSingleDestination(c echo.Context) error
	GetAllDestination(c echo.Context) error
}
