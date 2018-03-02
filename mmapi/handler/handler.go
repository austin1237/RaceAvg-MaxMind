package handler

import (
	"github.com/labstack/echo"
	"github.com/user/mmapi/location"
)

func GetLocations(c echo.Context) error {
	ipAddress := c.QueryParam("ip")
	location, err := location.GetLocationByIP(ipAddress)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(200, &location)
}

func GetHealth(c echo.Context) error {
	return c.String(200, "")
}
