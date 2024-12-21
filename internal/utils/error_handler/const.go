package errorhandler

import (
	"github.com/Nie-Mand/go-api/internal/core/services/hello"
)

var (
	ERRORS_MAP = map[error]int{
		hello.ErrHelloService: 401,
	}
)
