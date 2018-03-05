package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/user/mmapi/maxmind"
	"github.com/user/mmapi/route"
)

func init() {
	maxmind.MustLoadCSVs()
}

func CreateApi() *echo.Echo {
	api := echo.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	route.SetUpRoutes(api)
	return api
}

func main() {
	api := CreateApi()
	api.Logger.Fatal(api.Start(":8080"))
}
