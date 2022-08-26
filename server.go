package main

import (
	// MyMind "random_words/randoms"
	route "random_words/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	route.AllRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
