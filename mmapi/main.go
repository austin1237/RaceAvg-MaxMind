package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/user/mmapi/maxmind"
	"github.com/user/mmapi/route"
)

func init() {
	maxmind.LoadCSVs()
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	route.SetUpRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
