package randoms

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	random struct {
		ID   int    `json: "id"`
		Desc string `json: "desc"`
	}
)

var (
	randoms = map[int]*random{}
	seq     = 1
)

func GetRandoms(c echo.Context) error {
	return c.JSON(http.StatusOK, randoms)
}

func GetRandom(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, randoms[id])
}

func AddData(c echo.Context) error {
	u := &random{
		ID: seq,
	}

	if err := c.Bind(u); err != nil {
		return err
	}

	randoms[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func UpdateData(c echo.Context) error {
	u := new(random)

	if err := c.Bind(u); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	randoms[id].Desc = u.Desc

	return c.JSON(http.StatusOK, randoms[id])
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	delete(randoms, id)

	return c.NoContent(http.StatusNoContent)
}
