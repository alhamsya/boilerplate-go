package restMiddleware

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_resp"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (m *Middleware) Authorize(h fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//get the JWT string from the cookie
		tokenString := strings.ReplaceAll(ctx.GetReqHeaders()["Authorization"], "Bearer ", "")

		// Initialize a new instance of `Claims`
		claims := &MyClaims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.Cfg.JWT.SignatureKey), nil
		})

		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return customResp.Err(ctx, fiber.StatusUnauthorized, err, "ErrSignatureInvalid")
		}

		if err != nil {
			return customResp.Err(ctx, fiber.StatusBadRequest, err, "ParseWithClaims")
		}

		if !tkn.Valid {
			return customResp.Err(ctx, fiber.StatusUnauthorized, fmt.Errorf("invalid token"))
		}

		ctx.SetUserContext(context.WithValue(ctx.UserContext(), "user", *claims))
		return h(ctx)
	}
}
