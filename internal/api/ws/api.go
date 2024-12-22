package ws_api

import (
	"github.com/Nie-Mand/go-api/internal/api"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

// @tag.name			Websockets
// @tag.description	An Endpoint that handles websockets
type WSAPI struct {
	*api.API
}

func Register(api *api.API) {
	wsApi := &WSAPI{
		API: api,
	}

	group := api.E.Group("")
	group.GET("/ws", wsApi.handler)
}
