package hello_api

import (
	"github.com/Nie-Mand/go-api/internal/utils"
	errorhandler "github.com/Nie-Mand/go-api/internal/utils/error_handler"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary				Hello world
// @Description 			Hello world
// @securityDefinitions.ConfigAPIkey ConfigAPIKeyAuth
// @in header
// @name Authorization
// @Tags			Hello
// @Accept			json
// @Success		200	{object} utils.APIResponse{result=WorldResponse}
// @Router			/hello/world [get]
func (api *HelloController) World(c echo.Context) error {
	hello, err := api.R.Hello.HelloWorld()
	if err != nil {
		zap.L().Error("Failed to get hello world", zap.Error(err))
		return c.JSON(errorhandler.NewAPIError(api.R.Hello.FormatError(err)))
	}

	response := WorldResponse{
		Hello: hello.World,
	}

	return c.JSON(200, utils.NewAPIResponse("success", response))
}

type WorldResponse struct {
	Hello string `json:"hello"`
}
