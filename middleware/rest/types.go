package restMiddleware

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/golang-jwt/jwt/v4"
)

type Middleware struct {
	Cfg       *config.ServiceConfig
	DBRepo    repository.DBRepo
	UtilsRepo repository.UtilsRepo
}

// MyClaims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type MyClaims struct {
	jwt.StandardClaims
	Email    string `json:"email"`
	Password string `json:"password"`
}
