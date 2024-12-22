package authenticate

import (
	"errors"

	"github.com/Nie-Mand/go-api/internal/api"
	"github.com/Nie-Mand/go-api/internal/utils"
	errorhandler "github.com/Nie-Mand/go-api/internal/utils/error_handler"
	"github.com/labstack/echo/v4"
)

func NewAuthenticator(api *api.API) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := utils.GetTokenFromAuthorizationHeader(
				c.Request().Header.Get("Authorization"),
			)

			if token == "" {
				return c.JSON(errorhandler.NewAPIError(
					errors.New("error getting token from authorization header"),
					401,
				))
			}

			// Get the user from the token

			return next(c)
		}
	}
}
