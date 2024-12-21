package repositories

import (
	"github.com/Nie-Mand/go-api/internal/core/domain"
	"github.com/Nie-Mand/go-api/pkg/db"
)

type HelloRepository struct {
	DB *db.DB
}

func NewHelloRepository(db *db.DB) *HelloRepository {
	return &HelloRepository{
		DB: db,
	}
}

func (r *HelloRepository) GetHello() (domain.Hello, error) {
	return domain.Hello{World: "world"}, nil
}
