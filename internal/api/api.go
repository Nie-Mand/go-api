package api

import (
	// "github.com/Nie-Mand/anas-init/internal/core/services/auth"
	// "github.com/Nie-Mand/anas-init/internal/core/services/sites"
	"github.com/labstack/echo/v4"
)

type API struct {
	E *echo.Echo

	// Services
	// Auth  *auth.AuthService
	// Sites *sites.SitesService
}

type APICfg func(*API)

func NewAPI(cfg ...APICfg) *API {
	api := &API{}

	for _, c := range cfg {
		c(api)
	}

	return api
}
