package controller

import (
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/pkg/dto"
	"ecommerce-evermos-projects/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type CategoryControllerImpl struct {
	uc usecase.CategoryUseCase
}

func NewCategoryController(uc usecase.CategoryUseCase) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		uc: uc,
	}
}

func (ctl *CategoryControllerImpl) GetAllCategory(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, err := ctl.uc.GetAllCategory(c)

	if err != nil {
		return helper.ErrorResponse(ctx, err.Code, err.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, res)
}

func (ctl *CategoryControllerImpl) GetCategoryByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	id, err := ctx.ParamsInt("category_id")
	if err != nil || id == 0 {
		return helper.ErrorResponse(ctx, fiber.StatusNotFound, "record not found.")
	}

	res, errUC := ctl.uc.GetCategoryByID(c, uint(id))

	if errUC != nil {
		return helper.ErrorResponse(ctx, errUC.Code, err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, res)
}

func (ctl *CategoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(dto.CategoryReqCreate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}
	res, errUC := ctl.uc.CreateCategory(c, *data)
	if errUC != nil {
		return helper.ErrorResponse(ctx, errUC.Code, errUC.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusCreated, res)
}

func (ctl *CategoryControllerImpl) UpdateCategoryByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryId, err := ctx.ParamsInt("category_id")
	if err != nil || categoryId == 0 {
		return helper.ErrorResponse(ctx, fiber.StatusNotFound, err.Error())
	}
	data := new(dto.CategoryReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	res, errUC := ctl.uc.UpdateCategoryByID(c, uint(categoryId), *data)
	if errUC != nil {
		return helper.ErrorResponse(ctx, errUC.Code, errUC.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusCreated, res)
}

func (ctl *CategoryControllerImpl) DeleteCategoryByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryId, er := ctx.ParamsInt("category_id")
	if er != nil || categoryId == 0 {
		return helper.ErrorResponse(ctx, fiber.StatusNotFound, "record not found")
	}

	res, err := ctl.uc.DeleteCategoryByID(c, uint(categoryId))

	if err != nil {
		return helper.ErrorResponse(ctx, err.Code, err.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, res)
}
