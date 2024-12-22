package hello_api

import "github.com/Nie-Mand/go-api/internal/api"

// @tag.name			Hello
// @tag.description	An Endpoint that handles hello world
type HelloController struct {
	R *api.API
}

func RegisterHelloController(_api *api.API) {
	c := &HelloController{
		R: _api,
	}

	group := _api.E.Group("/hello")

	group.GET("/world", c.World)
}
