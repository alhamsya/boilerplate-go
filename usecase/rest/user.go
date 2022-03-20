package restUC

import (
	"fmt"
	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/domain/models/database"
	"github.com/alhamsya/boilerplate-go/domain/models/request"
	"github.com/alhamsya/boilerplate-go/domain/models/response"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"github.com/gofiber/fiber/v2"
)

func (uc *UCInteractor) DoCreateUser(ctx *fiber.Ctx, reqClient *modelReq.User) (resp *modelResp.User, httpCode int, err error) {
	respDB, err := uc.DBRepo.GetUserByEmail(ctx.Context(), reqClient.Email)
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "database", "GetUserByEmail")
	}

	if respDB != nil {
		return nil, fiber.StatusBadRequest, fmt.Errorf("email has been exist")
	}

	tx, err := uc.DBRepo.TxBegin(ctx.Context())
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "transaction", "TxBegin")
	}

	timeNow, err := uc.UtilsRepo.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "utils", "CurrentTimeF")
	}

	hashPassword, err := uc.UtilsRepo.HashPassword(reqClient.Password)
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "utils", "HashPassword")
	}

	reqDB := &modelDB.User{
		Email:     reqClient.Email,
		Password:  hashPassword,
		Status:    1,
		CreatedAt: timeNow,
		CreatedBy: reqClient.Email,
	}

	id, err := uc.DBRepo.CreateUserTx(ctx.Context(), tx, reqDB)
	if err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "database", "CreateUserTx")
	}

	if err := uc.DBRepo.TxCommit(tx); err != nil {
		return nil, fiber.StatusInternalServerError, customError.WrapFlag(err, "transaction", "TxCommit")
	}

	resp = &modelResp.User{
		ID:        id,
		Email:     reqDB.Email,
		Status:    reqDB.Status,
		CreatedAt: reqDB.CreatedAt,
		CreatedBy: reqDB.CreatedBy,
	}
	return resp, fiber.StatusCreated, nil
}
