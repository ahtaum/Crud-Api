package routes

import (
	MyMind "random_words/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", MyMind.GetRandoms)
	e.GET("/random/:id", MyMind.GetRandom)
	e.POST("/add", MyMind.AddData)
	e.PUT("/update/:id", MyMind.UpdateData)
	e.DELETE("/delete/:id", MyMind.DeleteUser)

	return e
}
