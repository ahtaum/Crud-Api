package routes

import (
	MyMind "random_words/randoms"

	"github.com/labstack/echo/v4"
)

func AllRoutes(e *echo.Echo) {
	e.GET("/", MyMind.GetRandoms)
	e.GET("/random/:id", MyMind.GetRandom)
	e.POST("/add", MyMind.AddData)
	e.PUT("/update/:id", MyMind.UpdateData)
	e.DELETE("/delete/:id", MyMind.DeleteUser)
}
