package customResp

import "github.com/gofiber/fiber/v2"

type Response struct {
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	error    error
	httpCode int
	ctx      *fiber.Ctx
}
