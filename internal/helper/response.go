package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func SuccessResponse(ctx *fiber.Ctx, code int, message string, data interface{}) error {
	return ctx.Status(code).JSON(Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, code int, message string, errors interface{}) error {
	return ctx.Status(code).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
	})
}
