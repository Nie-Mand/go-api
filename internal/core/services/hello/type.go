package hello

import (
	"github.com/Nie-Mand/go-api/internal/core/ports"
)

type HelloService struct {
	Hello  ports.HelloRepository
	PubSub ports.PubSub
}

type HelloServiceCfg func(a *HelloService)
