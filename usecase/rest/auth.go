package restUC

import (
	"fmt"
	"time"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/domain/models/database"
	"github.com/alhamsya/boilerplate-go/domain/models/request"
	"github.com/alhamsya/boilerplate-go/domain/models/response"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	restMiddleware "github.com/alhamsya/boilerplate-go/middleware/rest"
)

func (uc *UCInteractor) DoSignup(ctx *fiber.Ctx, reqClient *modelReq.User) (resp *modelResp.User, httpCode int, err error) {
	respDB, err := uc.DBRepo.GetUserByEmail(ctx.Context(), reqClient.Email)
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "database", "GetUserByEmail")
	}

	if respDB != nil {
		return nil, fiber.StatusBadRequest, fmt.Errorf("email has been exist")
	}

	timeNow, err := uc.UtilsRepo.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "utils", "CurrentTimeF")
	}

	hashPassword, err := uc.UtilsRepo.HashPassword(reqClient.Password)
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "utils", "HashPassword")
	}

	claims := restMiddleware.MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(uc.Cfg.JWT.ExpDurationInHours) * time.Hour).Unix(),
		},
		Email: reqClient.Email,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(uc.Cfg.JWT.SignatureKey))
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "jwt", "SignedString")
	}

	tx, err := uc.DBRepo.TxBegin(ctx.Context())
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "transaction", "TxBegin")
	}

	reqDB := &modelDB.User{
		Email:     reqClient.Email,
		Password:  hashPassword,
		Status:    1,
		CreatedAt: timeNow,
		CreatedBy: reqClient.Email,
	}

	id, err := uc.DBRepo.CreateUserTx(ctx.Context(), tx, reqDB)
	if err := uc.DBRepo.TxError(tx, err); err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "database", "CreateUserTx")
	}

	if err := uc.DBRepo.TxCommit(tx); err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "transaction", "TxCommit")
	}

	resp = &modelResp.User{
		ID:        id,
		Email:     reqDB.Email,
		Token:     signedToken,
		Status:    reqDB.Status,
		CreatedAt: reqDB.CreatedAt,
		CreatedBy: reqDB.CreatedBy,
	}
	return resp, fiber.StatusCreated, nil
}

func (uc *UCInteractor) DoSigning(ctx *fiber.Ctx, reqClient *modelReq.User) (resp *modelResp.User, httpCode int, err error) {
	respDB, err := uc.DBRepo.GetUserByEmail(ctx.Context(), reqClient.Email)
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "database", "GetUserByEmail")
	}

	if respDB == nil {
		return nil, fiber.StatusNotFound, fmt.Errorf("data user email invalid")
	}

	if !uc.UtilsRepo.CheckPasswordHash(reqClient.Password, respDB.Password) {
		return nil, fiber.StatusUnauthorized, fmt.Errorf("data user password invalid")
	}

	timeNow, err := uc.UtilsRepo.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "utils", "CurrentTimeF")
	}

	// Create the JWT claims, which includes the username and expiry time
	claims := restMiddleware.MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(uc.Cfg.JWT.ExpDurationInHours) * time.Hour).Unix(),
		},
		Email: reqClient.Email,
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(uc.Cfg.JWT.SignatureKey))
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "jwt", "SignedString")
	}

	reqDB := &modelDB.Signing{
		Email:     reqClient.Email,
		CreatedAt: timeNow,
		CreatedBy: respDB.Email,
	}
	_, err = uc.DBRepo.CreateHistorySigning(ctx.Context(), reqDB)
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "database", "CreateHistorySigning")
	}

	resp = &modelResp.User{
		ID:        respDB.ID,
		Email:     respDB.Email,
		Token:     signedToken,
		Status:    respDB.Status,
		CreatedAt: respDB.CreatedAt,
		CreatedBy: respDB.CreatedBy,
	}

	return resp, fiber.StatusOK, nil
}
