package internal

import (
	"github.com/Nie-Mand/go-api/internal/api"
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

func Run() {
	// db, err := db.New()
	// handleError(err)
	// defer db.Close()

	// usersR := repos.NewUsersRepository(db)
	// sitesR := repos.NewSitesRepository(db)

	e := echo.NewEchoServer(echo.NewEchoConfig("8080"))

	// Services
	// authS := auth.NewAuthService(func(a *auth.AuthService) {
	// 	a.Users = usersR
	// 	a.Bcrypt = bcrypt.NewBcrypt()
	// 	a.JWT = jwt.NewJWT("secret-string", 10)
	// })

	// sitesS := sites.NewSitesService(func(s *sites.SitesService) {
	// 	s.Sites = sitesR
	// })

	// Register API
	api.NewAPI(func(a *api.API) {
		a.E = e.E
		// a.Auth = authS
		// a.Sites = sitesS
	})
	// auth_api.RegisterAuthController(_api)
	// sites_api.RegisterSitesController(_api)
	e.Start()
}
