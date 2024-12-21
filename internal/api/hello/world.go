package hello_api

import (
	"github.com/Nie-Mand/go-api/internal/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (api *HelloController) World(c echo.Context) error {

	hello, err := api.R.Hello.HelloWorld()
	if err != nil {
		zap.L().Error("Failed to get hello world", zap.Error(err))
		return c.JSON(utils.NewAPIErrorResponse(400, api.R.Hello.FormatError(err)))
	}

	response := WorldResponse{
		Hello: hello.World,
	}

	return c.JSON(200, utils.NewAPIResponse[WorldResponse]("success", response))
}

type WorldResponse struct {
	Hello string `json:"hello"`
}
