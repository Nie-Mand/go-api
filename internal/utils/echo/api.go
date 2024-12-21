package echo

import (
	"github.com/Nie-Mand/go-api/internal/utils/middlewares"
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

	// Setup the middlewares here
	e.Validator = middlewares.NewValidator()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return EchoServer{
		E:    e,
		Port: cfg.Port,
	}
}

func (e EchoServer) Start() error {
	zap.L().Info("Starting server on port " + e.Port)
	return e.E.Start(":" + e.Port)
}
