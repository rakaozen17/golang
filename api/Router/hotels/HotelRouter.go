package hotelsRouter

import (
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, s Hotels, m ...echo.MiddlewareFunc) {
	users := e.Group("v1/hotels")
	users.POST("/create", s.CreateHotel)
	users.PUT("/update", s.UpdateHotel)
	users.DELETE("/delete", s.DeleteHotel)
	users.GET("/:id", s.GetSingleHotel)
	users.GET("/", s.GetAllHotel)
}
