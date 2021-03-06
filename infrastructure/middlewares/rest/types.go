package restMiddleware

import (
	"github.com/alhamsya/boilerplate-go/domain/repositories"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"github.com/golang-jwt/jwt/v4"
)

type Middleware struct {
	Cfg       *config.ServiceConfig
	DBRepo    repositories.DBRepo
	UtilsRepo repositories.UtilsRepo
}

// MyClaims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type MyClaims struct {
	jwt.StandardClaims
	Email    string `json:"email"`
	Password string `json:"password"`
}
