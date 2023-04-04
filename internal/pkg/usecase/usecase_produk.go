package usecase

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/infrastructure/container"
	"ecommerce-evermos-projects/internal/pkg/dto"
	"ecommerce-evermos-projects/internal/pkg/repository"
	"ecommerce-evermos-projects/internal/pkg/validator"
	"ecommerce-evermos-projects/internal/utils/uploader"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type ProdukUseCase interface {
}

type ProdukUseCaseImpl struct {
	produkrepository repository.ProdukRepository
	apps             container.Apps
	container        container.Container
}

func NewProdukUseCase(produkrepository repository.ProdukRepository, apps container.Apps) ProdukUseCase {
	return &ProdukUseCaseImpl{
		produkrepository: produkrepository,
		apps:             apps,
	}

}

func (alc *ProdukUseCaseImpl) GetAllProduk(ctx context.Context, params dto.FilterProduk) (res []*dto.ProdukResp, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := alc.produkrepository.GetAllProduks(ctx, daos.FilterProduk{
		Limit:  params.Limit,
		Offset: params.Page,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("no data produk"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllProduk : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = dto.ProdukSliceDaosToDto(resRepo)

	return res, nil
}

func (alc *ProdukUseCaseImpl) GetProdukByID(ctx context.Context, id uint) (res dto.ProdukResp, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.produkrepository.GetProdukByID(ctx, id)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("no data produk"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = dto.ProdukDaosToDto(resRepo)

	return res, nil
}

func (alc *ProdukUseCaseImpl) CreateProduk(ctx context.Context, data dto.ProdukReqCreate) (res uint, err *helper.ErrorStruct) {
	var validationErrors error

	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	// upload each photo
	for _, file := range data.Photo {
		er := validator.ValidateFile(file, 2*1024*1024, []string{".jpg", "jpeg", ".png"})
		if er != nil {
			validationErrors = errors.Join(er)
		}
		fileName := strings.Split(file.Filename, ".")
		name := fmt.Sprintf("produk_%s-%d", fileName[0], time.Now().Unix())
		image := fmt.Sprintf("%s.%s", name, fileName[1])
		uploader.SaveFile(file, fmt.Sprintf("/images/%s", image))
	}

	if validationErrors != nil {
		return res, &helper.ErrorStruct{
			Err:  validationErrors,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.produkrepository.CreateProduk(ctx, daos.Produk{
		Nama:          data.Nama,
		Slug:          slug.Make(data.Nama + " " + strconv.Itoa(rand.Int())),
		HargaReseller: data.HargaReseller,
		HargaKonsumen: data.HargaKonsumen,
		Stok:          data.Stok,
		Deskripsi:     data.Deskripsi,
		CategoryID:    data.CategoryID,
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

func (alc *ProdukUseCaseImpl) UpdateProdukByID(ctx context.Context, id uint, data dto.ProdukReqUpdate) (res string, err *helper.ErrorStruct) {
	var validationErrors error
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}
	produk := daos.Produk{
		Nama:          data.Nama,
		Slug:          slug.Make(data.Nama + " " + strconv.Itoa(rand.Int())),
		HargaReseller: data.HargaReseller,
		HargaKonsumen: data.HargaKonsumen,
		Stok:          data.Stok,
		Deskripsi:     data.Deskripsi,
		CategoryID:    data.CategoryID,
	}

	// upload each photo
	for _, file := range data.Photo {
		er := validator.ValidateFile(file, 2*1024*1024, []string{".jpg", "jpeg", ".png"})
		if er != nil {
			validationErrors = errors.Join(er)
		}
		fotoUrl, er := alc.container.Storage.UploadFile(file, "produk")
		if er != nil {
			return res, &helper.ErrorStruct{
				Err:  er,
				Code: fiber.StatusInternalServerError,
			}
		}

		produk.FotoProduks = append(produk.FotoProduks, daos.FotoProduk{
			ProdukID: id,
			Url:      fotoUrl,
		})

	}

	if validationErrors != nil {
		return res, &helper.ErrorStruct{
			Err:  validationErrors,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.produkrepository.UpdateProdukByID(ctx, id, produk)

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at UpdateProdukByID : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (alc *ProdukUseCaseImpl) DeleteProdukByID(ctx context.Context, id uint) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.produkrepository.DeleteProdukByID(ctx, id)

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at DeleteProdukByID : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
