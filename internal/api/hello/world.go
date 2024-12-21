package hello_api

import (
	"github.com/Nie-Mand/go-api/internal/utils"
	"github.com/labstack/echo/v4"
)

func (api *HelloController) World(c echo.Context) error {
	response := WorldResponse{
		Hello: "world",
	}

	return c.JSON(200, utils.NewAPIResponse[WorldResponse]("success", response))
}

type WorldResponse struct {
	Hello string `json:"hello"`
}
