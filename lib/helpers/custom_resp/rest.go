package customResp

import (
	"strings"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/friendsofgo/errors"
	"github.com/gofiber/fiber/v2"
)

func Err(ctx *fiber.Ctx, httpCode int, err error, arg ...string) (resp error) {
	args := strings.Join(arg, "|")
	switch httpCode / 100 {
	case fiber.StatusBadRequest / 100:
		if strings.TrimSpace(args) == "" {
			args = errors.Cause(err).Error()
		}
		resp = ctx.Status(httpCode).JSON(&fiber.Map{
			"message": args,
			"data":    nil,
		})
	case fiber.StatusInternalServerError / 100:
		customLog.ErrorF("[%s][USECASE] %s: %v", ctx.OriginalURL(), args, err)
		resp = ctx.Status(httpCode).JSON(&fiber.Map{
			"message": "please try again",
			"data":    nil,
		})
	}

	return resp
}

func Success(ctx *fiber.Ctx, httpCode int, data interface{}, message string) error {
	return ctx.Status(httpCode).JSON(&fiber.Map{
		"message": message + " successfully",
		"data":    data,
	})
}
