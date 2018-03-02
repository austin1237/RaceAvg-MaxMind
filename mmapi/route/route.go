package route

import (
	"github.com/labstack/echo"
	"github.com/user/mmapi/handler"
)

func SetUpRoutes(e *echo.Echo) {
	e.GET("/location", handler.GetLocations)
	e.GET("/health", handler.GetHealth)
}
