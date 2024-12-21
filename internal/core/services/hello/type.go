package hello

import (
	"errors"

	"github.com/Nie-Mand/go-api/internal/core/ports"
)

type HelloService struct {
	Hello  ports.HelloRepository
	PubSub ports.PubSub
}

type HelloServiceCfg func(a *HelloService)

func NewHelloService(cfg ...HelloServiceCfg) *HelloService {
	s := &HelloService{}

	for _, c := range cfg {
		c(s)
	}

	return s
}

func (s *HelloService) FormatError(err error) error {
	switch err {
	case ErrHelloService:
		return err

	default:
		return errors.New("internal server error")
	}
}
