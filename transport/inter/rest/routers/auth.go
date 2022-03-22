package restRouters

import (
	"errors"
	"strings"
	"time"

	"github.com/alhamsya/boilerplate-go/domain/models/request"
	"github.com/alhamsya/boilerplate-go/domain/models/response"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_resp"
	"github.com/alhamsya/boilerplate-go/middleware/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (rest *RestServer) CreateRefreshToken(ctx *fiber.Ctx) error {
	tokenString := strings.ReplaceAll(ctx.GetReqHeaders()["Authorization"], "Bearer ", "")

	claims := &restMiddleware.MyClaims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(rest.Cfg.JWT.SignatureKey), nil
	})

	if errors.Is(err, jwt.ErrSignatureInvalid) {
		return customResp.New(ctx).SetHttpCode(fiber.StatusUnauthorized).SetErr(err).Send("invalid signature")
	}

	if err != nil {
		return customResp.New(ctx).SetHttpCode(fiber.StatusBadRequest).SetErr(err).Send("invalid claim authentication")
	}

	if !tkn.Valid {
		return customResp.New(ctx).SetHttpCode(fiber.StatusUnauthorized).Send("invalid token")
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 1 minutes of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > time.Duration(rest.Cfg.JWT.ElapsedDurationInMinutes)*time.Minute {
		return customResp.New(ctx).SetHttpCode(fiber.StatusBadRequest).Send("invalid time expiry")
	}

	// Now, create a new token for the current use, with a renewed expiration time
	claims.ExpiresAt = time.Now().Add(time.Duration(rest.Cfg.JWT.ExpDurationInHours) * time.Hour).Unix()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(rest.Cfg.JWT.SignatureKey))
	if err != nil {
		return customResp.New(ctx).SetErr(err).SetHttpCode(fiber.StatusInternalServerError).Send("SignedString")
	}

	resp := modelResp.Auth{
		Token: signedToken,
	}

	return customResp.New(ctx).SetHttpCode(fiber.StatusOK).SetData(resp).SetMessage("create refresh token").Send()
}

func (rest *RestServer) CreateSignup(ctx *fiber.Ctx) error {
	req := new(modelReq.User)
	err := ctx.BodyParser(req)
	if err != nil {
		return customResp.New(ctx).SetErr(err).SetHttpCode(fiber.StatusBadRequest).Send()
	}

	resp, httpCode, err := rest.RestInteractor.Usecase.DoSignup(ctx, req)
	if err != nil {
		return customResp.New(ctx).SetErr(err).SetHttpCode(httpCode).Send()
	}

	return customResp.New(ctx).SetHttpCode(httpCode).SetData(resp).SetMessage("create signup").Send()
}

func (rest *RestServer) CreateSigning(ctx *fiber.Ctx) error {
	req := new(modelReq.User)
	err := ctx.BodyParser(req)
	if err != nil {
		return customResp.New(ctx).SetErr(err).SetHttpCode(fiber.StatusBadRequest).Send()
	}

	resp, httpCode, err := rest.RestInteractor.Usecase.DoSigning(ctx, req)
	if err != nil {
		return customResp.New(ctx).SetErr(err).SetHttpCode(httpCode).Send()
	}

	return customResp.New(ctx).SetHttpCode(httpCode).SetData(resp).SetMessage("create signing").Send()
}
