package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/user/mmapi/adapter"
	"github.com/user/mmapi/maxmind"
	"github.com/user/mmapi/validator"
)

func GetLocations(c echo.Context) error {
	clientIP := c.QueryParam("ip")
	clientIP = strings.Replace(clientIP, "\"", "", -1)
	clientIP = strings.Replace(clientIP, "'", "", -1)
	err := validator.ValidateLocationQuery(clientIP)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbIP, err := adapter.ClientIPToDb(clientIP)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbLocation, err := maxmind.QueryForLocation(dbIP)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	clientLocation := adapter.DbLocationToClient(dbLocation)
	return c.JSON(http.StatusOK, &clientLocation)
}

func GetHealth(c echo.Context) error {
	return c.String(http.StatusOK, "")
}
