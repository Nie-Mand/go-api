package api

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Validate[B any](c echo.Context) (*B, error) {
	req := new(B)
	if err := c.Bind(req); err != nil {
		zap.L().Error("Failed to bind request", zap.Error(err))
		return nil, err
	}

	if err := c.Validate(&req); err != nil {
		zap.L().Error("Failed to validate request", zap.Error(err))
		return nil, err
	}

	return req, nil
}
