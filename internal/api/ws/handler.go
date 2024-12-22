package ws_api

import (
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

// @Summary		Handle socket
// @Description 			Handle socket
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Tags			Websockets
// @Router			/ws [get]
func (a *WSAPI) handler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			zap.L().Error("error reading message", zap.Error(err))
			break
		}

		zap.L().Info("received message", zap.String("message", string(msg)))
		err = ws.WriteJSON(msg)
		if err != nil {
			zap.L().Error("error writing message", zap.Error(err))
			break
		}
	}

	return nil
}
