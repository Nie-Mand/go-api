package api

import (
	"github.com/Nie-Mand/go-api/internal/core/services/hello"
	"github.com/labstack/echo/v4"
)

type API struct {
	E *echo.Echo

	// Services
	Hello *hello.HelloService
}

type APICfg func(*API)

func NewAPI(cfg ...APICfg) *API {
	api := &API{}

	for _, c := range cfg {
		c(api)
	}

	return api
}
