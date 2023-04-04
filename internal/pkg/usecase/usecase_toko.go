package usecase

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/infrastructure/storage"
	"ecommerce-evermos-projects/internal/pkg/dto"
	"ecommerce-evermos-projects/internal/pkg/repository"
	"ecommerce-evermos-projects/internal/pkg/validator"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TokoUseCase interface {
	GetMyToko(ctx context.Context, user daos.User) (toko dto.TokoResp, err *helper.ErrorStruct)
	GetTokoByID(ctx context.Context, id uint) (toko dto.TokoResp, err *helper.ErrorStruct)
	FetchToko(ctx context.Context, params dto.FilterToko) (res []dto.TokoResp, err *helper.ErrorStruct)
	UpdateToko(ctx context.Context, id uint, user daos.User, data dto.TokoReqUpdate) (res string, err *helper.ErrorStruct)
}

type TokoUseCaseImpl struct {
	repo    repository.TokoRepository
	storage storage.Storage
}

func NewTokoUseCase(repo repository.TokoRepository) TokoUseCase {
	return &TokoUseCaseImpl{
		repo: repo,
	}
}

func (uc *TokoUseCaseImpl) GetMyToko(ctx context.Context, user daos.User) (toko dto.TokoResp, err *helper.ErrorStruct) {
	resRepo, errRepo := uc.repo.FindTokoByUserID(ctx, user.ID)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return toko, &helper.ErrorStruct{
			Err:  errors.New("record not found"),
			Code: fiber.StatusNotFound,
		}
	}

	if err != nil {
		return toko, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}
	toko = dto.TokoResp{
		ID:       resRepo.ID,
		Nama:     resRepo.Name,
		PhotoUrl: resRepo.PhotoUrl,
		UserID:   resRepo.UserID,
	}

	return toko, nil
}

func (uc *TokoUseCaseImpl) GetTokoByID(ctx context.Context, id uint) (toko dto.TokoResp, err *helper.ErrorStruct) {
	resRepo, errRepo := uc.repo.FindTokoByID(ctx, id)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return toko, &helper.ErrorStruct{
			Err:  errors.New("Toko tidak ditemukan"),
			Code: fiber.StatusNotFound,
		}
	}

	if err != nil {
		return toko, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}
	toko = dto.TokoResp{
		ID:       resRepo.ID,
		Nama:     resRepo.Name,
		PhotoUrl: resRepo.PhotoUrl,
	}

	return toko, nil
}

func (uc *TokoUseCaseImpl) FetchToko(ctx context.Context, params dto.FilterToko) (res []dto.TokoResp, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := uc.repo.GetToko(ctx, daos.FilterToko{
		Name:   params.Nama,
		Limit:  params.Limit,
		Offset: params.Page,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Err:  errors.New("Toko tidak ditemukan"),
			Code: fiber.StatusNotFound,
		}
	}

	if err != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}
	for _, v := range resRepo {
		res = append(res, dto.TokoResp{
			ID:       v.ID,
			Nama:     v.Name,
			PhotoUrl: v.PhotoUrl,
		})
	}

	return res, nil
}

func (uc *TokoUseCaseImpl) UpdateToko(ctx context.Context, id uint, user daos.User, data dto.TokoReqUpdate) (res string, err *helper.ErrorStruct) {
	er := validator.ValidateFile(data.Photo, 2*1024*1024, []string{".jpg", "jpeg", ".png"})
	if er != nil {
		return res, &helper.ErrorStruct{
			Err:  er,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := uc.repo.FindTokoByUserID(ctx, user.ID)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Err:  errors.New("record not found"),
			Code: fiber.StatusNotFound,
		}
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}
	if resRepo.UserID != user.ID {
		return res, &helper.ErrorStruct{
			Err:  errors.New("forbidden"),
			Code: fiber.StatusForbidden,
		}
	}

	// save image to ./images dir
	fotoUrl, er := uc.storage.UploadFile(data.Photo, "toko")

	if er != nil {
		return res, &helper.ErrorStruct{
			Err:  er,
			Code: fiber.StatusInternalServerError,
		}
	}

	res, errRepo = uc.repo.UpdateToko(ctx, id, daos.Toko{
		Name:     data.Nama,
		PhotoUrl: &fotoUrl,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return res, nil
}
