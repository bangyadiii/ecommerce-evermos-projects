package helper

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func SuccessResponse(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(Response{
		Status:  true,
		Message: "Succeed to " + ctx.Method() + " data",
		Data:    data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, code int, errors string) error {
	var splittedError []string

	if len(errors) > 0 {
		splittedError = strings.Split(errors, "\n")
	}
	return ctx.Status(code).JSON(Response{
		Status:  false,
		Message: "Failed to " + ctx.Method() + " data",
		Errors:  splittedError,
	})
}
