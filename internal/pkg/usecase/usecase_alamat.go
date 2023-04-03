package usecase

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/pkg/dto"
	"ecommerce-evermos-projects/internal/pkg/repository"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AlamatUseCase interface {
	GetAllAlamat(ctx context.Context, user daos.User) (res []*dto.AlamatRes, err *helper.ErrorStruct)
	GetAlamatByID(ctx context.Context, user daos.User, alamatID uint) (res dto.AlamatRes, err *helper.ErrorStruct)
	CreateAlamat(ctx context.Context, user daos.User, data dto.AlamatReqCreate) (res uint, err *helper.ErrorStruct)
	UpdateAlamatByID(ctx context.Context, user daos.User, alamatID uint, data dto.AlamatReqUpdate) (res string, err *helper.ErrorStruct)
	DeleteAlamatByID(ctx context.Context, user daos.User, alamatID uint) (res string, err *helper.ErrorStruct)
}

type alamatUseCaseImpl struct {
	alamatRepo repository.AlamatRepository
}

func NewAlamatUseCase(repo repository.AlamatRepository) AlamatUseCase {
	return &alamatUseCaseImpl{alamatRepo: repo}
}

func (s *alamatUseCaseImpl) GetAllAlamat(ctx context.Context, user daos.User) (res []*dto.AlamatRes, err *helper.ErrorStruct) {
	resRepo, errRepo := s.alamatRepo.FindAllAlamat(ctx, user)
	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllAlamat: %s", errRepo.Error()))
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = dto.AlamatSliceDaosToDto(resRepo)
	return res, nil
}

func (s *alamatUseCaseImpl) GetAlamatByID(ctx context.Context, user daos.User, alamatID uint) (res dto.AlamatRes, err *helper.ErrorStruct) {
	resRepo, errRepo := s.alamatRepo.FindAlamatByID(ctx, alamatID, user)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errRepo,
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAlamatByID: %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = dto.AlamatDaosToDto(resRepo)
	return res, nil
}

func (alc *alamatUseCaseImpl) CreateAlamat(ctx context.Context, user daos.User, data dto.AlamatReqCreate) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateAlamat: %s", errValidate.Error()))
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}
	data.UserID = user.ID
	alamat := daos.Alamat{
		UserID:       data.UserID,
		Judul:        data.Judul,
		NamaPenerima: data.NamaPenerima,
		NoTelp:       data.NoTelp,
		DetailAlamat: data.DetailAlamat,
	}
	resRepo, errRepo := alc.alamatRepo.CreateAlamat(ctx, alamat)
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resRepo, nil
}

func (alc *alamatUseCaseImpl) UpdateAlamatByID(ctx context.Context, user daos.User, alamatID uint, data dto.AlamatReqUpdate) (res string, err *helper.ErrorStruct) {
	_, errRepo := alc.alamatRepo.FindAlamatByID(ctx, alamatID, user)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errRepo,
		}
	}

	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	errRepo = alc.alamatRepo.UpdateAlamat(ctx, daos.Alamat{
		Model: gorm.Model{
			ID: alamatID,
		},
		UserID:       user.ID,
		Judul:        data.Judul,
		NamaPenerima: data.NamaPenerima,
		NoTelp:       data.NoTelp,
		DetailAlamat: data.DetailAlamat,
	})

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return "", nil
}

func (alc *alamatUseCaseImpl) DeleteAlamatByID(ctx context.Context, user daos.User, alamatID uint) (res string, err *helper.ErrorStruct) {
	_, errRepo := alc.alamatRepo.FindAlamatByID(ctx, alamatID, user)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errRepo,
		}
	}

	errRepo = alc.alamatRepo.DeleteAlamat(ctx, alamatID)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return res, nil
}
