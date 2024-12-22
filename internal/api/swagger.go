package api

import (
	_ "github.com/Nie-Mand/go-api/internal/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Go API
//	@version		0.1
//	@description	This backend behind the Go API

//	@BasePath	/

// @securitydefinitions.jwt
// @in	header
// @name	Authorization
func RegisterSwaggerDocumentation(_api *API) {
	_api.E.GET("/docs/*", echoSwagger.WrapHandler)
}
