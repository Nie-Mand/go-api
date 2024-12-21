package internal

import (
	"github.com/Nie-Mand/go-api/internal/api"
	hello_api "github.com/Nie-Mand/go-api/internal/api/hello"
	"github.com/Nie-Mand/go-api/internal/core/services/hello"
	"github.com/Nie-Mand/go-api/pkg/db"
	"github.com/Nie-Mand/go-api/pkg/db/repositories"
	"github.com/Nie-Mand/go-api/pkg/gateways/pubsub"

	// auth_api "github.com/Nie-Mand/anas-init/internal/api/auth"
	// sites_api "github.com/Nie-Mand/anas-init/internal/api/sites"
	// "github.com/Nie-Mand/anas-init/internal/core/services/auth"
	// "github.com/Nie-Mand/anas-init/internal/core/services/sites"
	"github.com/Nie-Mand/go-api/internal/utils/echo"
	// "github.com/Nie-Mand/anas-init/pkg/db"
	// repos "github.com/Nie-Mand/anas-init/pkg/db/repositories"
	// "github.com/Nie-Mand/anas-init/pkg/utils/bcrypt"
	// "github.com/Nie-Mand/anas-init/pkg/utils/jwt"
)

func Run() error {
	db, err := db.New()
	if err != nil {
		return err
	}

	defer db.Close()

	// Repositories and Gateways
	helloR := repositories.NewHelloRepository(db)
	ps := pubsub.NewPubSubClient()

	// Services
	helloS := hello.NewHelloService(func(s *hello.HelloService) {
		s.Hello = helloR
		s.PubSub = ps
	})

	// Register API
	e := echo.NewEchoServer(echo.NewEchoConfig("8080"))
	_api := api.NewAPI(func(a *api.API) {
		a.E = e.E
		a.Hello = helloS
	})
	hello_api.RegisterHelloController(_api)
	return e.Start()
}
