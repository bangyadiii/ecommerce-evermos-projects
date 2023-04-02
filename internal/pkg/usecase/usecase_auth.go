package usecase

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/infrastructure/container"
	"ecommerce-evermos-projects/internal/pkg/dto"
	"ecommerce-evermos-projects/internal/pkg/repository"
	"ecommerce-evermos-projects/internal/utils/jwt"
	"ecommerce-evermos-projects/internal/utils/password"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UsersUseCase interface {
	Login(ctx context.Context, params dto.UserReqLogin) (res dto.UserResLogin, err *helper.ErrorStruct)
	Register(ctx context.Context, params dto.UserReqRegister) (string, *helper.ErrorStruct)
	GetUser(ctx context.Context, email string) (res daos.User, err *helper.ErrorStruct)
}

type UsersUseCaseImpl struct {
	usersRepository repository.UsersRepository
	apps            container.Apps
}

func NewUsersUseCase(usersrepository repository.UsersRepository, apps container.Apps) UsersUseCase {
	return &UsersUseCaseImpl{
		usersRepository: usersrepository,
		apps:            apps,
	}

}

func (alc *UsersUseCaseImpl) Login(ctx context.Context, params dto.UserReqLogin) (res dto.UserResLogin, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.usersRepository.FindByNoTelp(ctx, params.NoTelp)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("user not found"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	errPass := password.ComparePassword([]byte(resRepo.Password), []byte(params.Password))
	if errPass != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errPass,
		}
	}
	log.Println("secret :", alc.apps.SecretJwt)
	log.Println("port :", alc.apps.HttpPort)
	token, errJWT := jwt.CreateJWT(int(resRepo.ID), resRepo.Email, []byte(alc.apps.SecretJwt))
	if errJWT != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  errPass,
		}
	}

	res = dto.UserResLogin{
		Nama:         resRepo.Nama,
		Email:        resRepo.Email,
		NoTelp:       resRepo.NoTelp,
		TanggalLahir: resRepo.TanggalLahir,
		Tentang:      resRepo.Tentang,
		Pekerjaan:    resRepo.Pekerjaan,
		Token:        token,
	}

	return res, nil
}

func (r *UsersUseCaseImpl) Register(ctx context.Context, params dto.UserReqRegister) (res string, err *helper.ErrorStruct) {
	_, errRepo := r.usersRepository.FindByEmail(ctx, params.Email)
	if !errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errors.New("email not available"),
		}
	}

	hashed, errPass := password.HashPassword([]byte(params.Password))
	if errPass != nil {
		return "", &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errPass,
		}
	}

	user := daos.User{
		Nama:         params.Nama,
		Email:        params.Email,
		NoTelp:       params.NoTelp,
		TanggalLahir: params.TanggalLahir,
		Tentang:      params.Tentang,
		Password:     string(hashed),
	}

	_, errRepo = r.usersRepository.SaveUser(ctx, user, daos.Toko{
		Name: user.Nama,
	})

	if err != nil {
		return "", &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return "Register Succeed", nil
}

func (alc *UsersUseCaseImpl) GetUser(ctx context.Context, email string) (res daos.User, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.usersRepository.FindByEmail(ctx, email)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("no data books"),
		}
	}
	if err != nil {
		return res, err
	}
	return resRepo, nil
}
