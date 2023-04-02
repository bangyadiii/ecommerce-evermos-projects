package usecase

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/pkg/dto"
	"ecommerce-evermos-projects/internal/pkg/repository"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryUseCase interface {
	GetAllCategory(ctx context.Context) (res []dto.CategoryRes, err *helper.ErrorStruct)
	GetCategoryByID(ctx context.Context, ID uint) (res dto.CategoryRes, err *helper.ErrorStruct)
	CreateCategory(ctx context.Context, data dto.CategoryReqCreate) (res uint, err *helper.ErrorStruct)
	UpdateCategoryByID(ctx context.Context, id uint, data dto.CategoryReqUpdate) (res string, err *helper.ErrorStruct)
	DeleteCategoryByID(ctx context.Context, id uint) (res string, err *helper.ErrorStruct)
}

type categoryUseCaseImpl struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUseCase(repo repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCaseImpl{categoryRepo: repo}
}

func (s *categoryUseCaseImpl) GetAllCategory(ctx context.Context) (res []dto.CategoryRes, err *helper.ErrorStruct) {
	resRepo, errRepo := s.categoryRepo.FindAllCategory(ctx)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errRepo,
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	res = dto.CategoriesDaosToDto(resRepo)

	return res, nil
}

func (s *categoryUseCaseImpl) GetCategoryByID(ctx context.Context, ID uint) (res dto.CategoryRes, err *helper.ErrorStruct) {
	resRepo, errRepo := s.categoryRepo.FindCategoryByID(ctx, ID)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errRepo,
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	res = dto.CategoryDaosToDto(resRepo)

	return res, nil
}

func (alc *categoryUseCaseImpl) CreateCategory(ctx context.Context, data dto.CategoryReqCreate) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.categoryRepo.CreateCategory(ctx, daos.Category{
		Nama: data.Nama,
	})

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (alc *categoryUseCaseImpl) UpdateCategoryByID(ctx context.Context, id uint, data dto.CategoryReqUpdate) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.categoryRepo.UpdateCategoryByID(ctx, id, daos.Category{
		Nama: data.Nama,
	})

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (alc *categoryUseCaseImpl) DeleteCategoryByID(ctx context.Context, id uint) (res string, err *helper.ErrorStruct) {
	_, errRepo := alc.categoryRepo.FindCategoryByID(ctx, id)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errRepo,
		}
	}
	res, errRepo = alc.categoryRepo.DeleteCategoryByID(ctx, id)

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return res, nil
}
