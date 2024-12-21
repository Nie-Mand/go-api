package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type EchoServer struct {
	E    *echo.Echo
	Port string
}

func NewEchoServer(cfg EchoConfig) EchoServer {
	e := echo.New()

	e.Validator = NewValidator()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return EchoServer{
		E:    e,
		Port: cfg.Port,
	}
}

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

func (e EchoServer) Start() {
	zap.L().Info("Starting server on port " + e.Port)
	err := e.E.Start(":" + e.Port)
	if err != nil {
		zap.L().Fatal(err.Error())
	}
}
