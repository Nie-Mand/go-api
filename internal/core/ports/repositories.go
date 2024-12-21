package ports

import "github.com/Nie-Mand/go-api/internal/core/domain"

type HelloRepository interface {
	GetHello() (domain.Hello, error)
}
