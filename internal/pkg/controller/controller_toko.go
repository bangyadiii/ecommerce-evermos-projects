package controller

import (
	"ecommerce-evermos-projects/internal/daos"
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/pkg/dto"
	"ecommerce-evermos-projects/internal/pkg/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TokoControllerImpl struct {
	uc usecase.TokoUseCase
}

func NewTokoController(uc usecase.TokoUseCase) *TokoControllerImpl {
	return &TokoControllerImpl{
		uc: uc,
	}
}

func (ctl *TokoControllerImpl) GetMyToko(ctx *fiber.Ctx) error {
	c := ctx.Context()

	// Retrieve the current_user value set by middleware
	currentUserVal, ok := ctx.Locals("current_user").(daos.User)
	if !ok {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve current user")
	}

	res, err := ctl.uc.GetMyToko(c, currentUserVal)
	if err != nil {
		return helper.ErrorResponse(ctx, err.Code, err.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, res)
}

func (ctl *TokoControllerImpl) GetTokoByID(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, er := ctx.ParamsInt("toko_id")
	if er != nil {
		return helper.ErrorResponse(ctx, fiber.StatusNotFound, "Not found")
	}

	res, err := ctl.uc.GetTokoByID(c, uint(id))
	if err != nil {
		return helper.ErrorResponse(ctx, err.Code, err.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, res)
}

func (ctl *TokoControllerImpl) FetchToko(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(dto.FilterToko)
	if err := ctx.QueryParser(filter); err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	res, err := ctl.uc.FetchToko(c, *filter)

	if err != nil {
		return helper.ErrorResponse(ctx, err.Code, err.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, fiber.Map{
		"page":  filter.Page,
		"limit": filter.Limit,
		"data":  res,
	})
}

func (ctl *TokoControllerImpl) UpdateToko(ctx *fiber.Ctx) error {
	c := ctx.Context()
	tokoID, err := ctx.ParamsInt("toko_id")

	if tokoID == 0 || err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusNotFound, "Not found")
	}

	data := new(dto.TokoReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, "Bad Request")
	}
	data.Photo, err = ctx.FormFile("photo")

	if err != nil {
		log.Println("image upload error --> ", err)
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, "server error")

	}

	// Retrieve the current_user value set by middleware
	currentUserVal, ok := ctx.Locals("current_user").(daos.User)
	if !ok {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve current user")
	}

	res, errUC := ctl.uc.UpdateToko(c, uint(tokoID), currentUserVal, *data)
	if errUC != nil {
		return helper.ErrorResponse(ctx, errUC.Code, errUC.Err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, res)
}
