package controller

import (
	"ecommerce-evermos-projects/internal/daos"
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/pkg/dto"
	"ecommerce-evermos-projects/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type AlamatControllerImpl struct {
	uc usecase.AlamatUseCase
}

func NewAlamatController(uc usecase.AlamatUseCase) *AlamatControllerImpl {
	return &AlamatControllerImpl{
		uc: uc,
	}
}

func (ctl *AlamatControllerImpl) GetAllAlamat(ctx *fiber.Ctx) error {
	c := ctx.Context()
	currentUserVal, ok := ctx.Locals("current_user").(daos.User)
	if !ok {
		// handle error
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve current user")
	}
	res, err := ctl.uc.GetAllAlamat(c, currentUserVal)

	if err != nil {
		return helper.ErrorResponse(ctx, err.Code, err.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, res)
}

func (ctl *AlamatControllerImpl) GetAlamatByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	id, err := ctx.ParamsInt("alamat_id")
	if err != nil || id == 0 {
		return helper.ErrorResponse(ctx, fiber.StatusNotFound, "record not found.")
	}
	// Retrieve the current_user value set by middleware
	currentUserVal, ok := ctx.Locals("current_user").(daos.User)
	if !ok {
		// handle error
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve current user")
	}
	res, errUC := ctl.uc.GetAlamatByID(c, currentUserVal, uint(id))

	if errUC != nil {
		return helper.ErrorResponse(ctx, errUC.Code, errUC.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, res)
}

func (ctl *AlamatControllerImpl) CreateAlamat(ctx *fiber.Ctx) error {
	c := ctx.Context()
	data := new(dto.AlamatReqCreate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}
	// Retrieve the current_user value set by middleware
	currentUserVal, ok := ctx.Locals("current_user").(daos.User)
	if !ok {
		// handle error
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve current user")
	}
	res, errUC := ctl.uc.CreateAlamat(c, currentUserVal, *data)
	if errUC != nil {
		return helper.ErrorResponse(ctx, errUC.Code, errUC.Err.Error())
	}
	return helper.SuccessResponse(ctx, fiber.StatusCreated, res)
}

func (ctl *AlamatControllerImpl) UpdateAlamatByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryId, err := ctx.ParamsInt("alamat_id")
	if err != nil || categoryId == 0 {
		return helper.ErrorResponse(ctx, fiber.StatusNotFound, err.Error())
	}
	data := new(dto.AlamatReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}
	// Retrieve the current_user value set by middleware
	currentUserVal, ok := ctx.Locals("current_user").(daos.User)
	if !ok {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve current user")
	}

	res, errUC := ctl.uc.UpdateAlamatByID(c, currentUserVal, uint(categoryId), *data)
	if errUC != nil {
		return helper.ErrorResponse(ctx, errUC.Code, errUC.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, res)
}

func (ctl *AlamatControllerImpl) DeleteAlamat(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryId, err := ctx.ParamsInt("alamat_id")
	if err != nil || categoryId == 0 {
		return helper.ErrorResponse(ctx, fiber.StatusNotFound, err.Error())
	}
	
	// Retrieve the current_user value set by middleware
	currentUserVal, ok := ctx.Locals("current_user").(daos.User)
	if !ok {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve current user")
	}

	res, errUC := ctl.uc.DeleteAlamatByID(c, currentUserVal, uint(categoryId))
	if errUC != nil {
		return helper.ErrorResponse(ctx, errUC.Code, errUC.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, res)
}

