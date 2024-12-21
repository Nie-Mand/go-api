package hello

import "github.com/Nie-Mand/go-api/internal/core/domain"

func (s *HelloService) HelloWorld() (domain.Hello, error) {
	err := s.PubSub.Pub("hello, world")
	if err != nil {
		return domain.Hello{}, ErrHelloService
	}

	return s.Hello.GetHello()
}
