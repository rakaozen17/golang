package destinationRouter

import (
	middleware "restproject/api/impl/middleware"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, s Destination, m ...echo.MiddlewareFunc) {
	users := e.Group("v1/destination")
	users.POST("/create", s.CreateDestination, middleware.ValidateToken)
	users.PUT("/update", s.UpdateDestination)
	users.DELETE("/delete", s.DeleteDestination)
	users.GET("/:id", s.GetSingleDestination)
	users.GET("/", s.GetAllDestination)
}
