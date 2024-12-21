package hello_api

import "github.com/Nie-Mand/go-api/internal/api"

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
