package controller

import (
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/pkg/dto"
	"ecommerce-evermos-projects/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	useruc usecase.UsersUseCase
}

func NewAuthControllerImpl(uc usecase.UsersUseCase) *AuthControllerImpl {
	return &AuthControllerImpl{uc}
}

func (uc *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	reqBody := new(dto.UserReqLogin)

	if err := ctx.BodyParser(reqBody); err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	res, err := uc.useruc.Login(ctx.Context(), *reqBody)
	if err != nil {
		return helper.ErrorResponse(ctx, err.Code, err.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK,  res)
}

func (uc *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	reqBody := new(dto.UserReqRegister)

	if err := ctx.BodyParser(reqBody); err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	res, err := uc.useruc.Register(ctx.Context(), *reqBody)

	if err != nil {
		return helper.ErrorResponse(ctx, err.Code, err.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK,  res)
}
