package handler

import (
	"ecommerce-evermos-projects/internal/infrastructure/container"
	"ecommerce-evermos-projects/internal/pkg/controller"
	"ecommerce-evermos-projects/internal/pkg/repository"
	"ecommerce-evermos-projects/internal/pkg/usecase"
	"ecommerce-evermos-projects/internal/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewAlamatRepository(containerConf.Mysqldb)
	userRepo := repository.NewUserRepository(containerConf.Mysqldb)
	uc := usecase.NewAlamatUseCase(repo)
	ucUser := usecase.NewUsersUseCase(userRepo, *containerConf.Apps)

	alamatCtl := controller.NewAlamatController(uc)
	authCtl := controller.NewAuthControllerImpl(ucUser)
	secret := containerConf.Apps.SecretJwt
	userAPI := r.Group("/user")
	userAPI.Get("/", middleware.VerifyToken(ucUser, secret), authCtl.GetCurrentUser)
	userAPI.Put("/", middleware.VerifyToken(ucUser, secret), authCtl.UpdateUser)

	alamatAPI := userAPI.Group("/alamat")

	alamatAPI.Get("/", middleware.VerifyToken(ucUser, secret), alamatCtl.GetAllAlamat)
	alamatAPI.Get("/:alamat_id", middleware.VerifyToken(ucUser, secret), alamatCtl.GetAlamatByID)
	alamatAPI.Post("/", middleware.VerifyToken(ucUser, secret), alamatCtl.CreateAlamat)
	alamatAPI.Put("/:alamat_id", middleware.VerifyToken(ucUser, secret), alamatCtl.UpdateAlamatByID)
	alamatAPI.Delete("/:alamat_id", middleware.VerifyToken(ucUser, secret), alamatCtl.DeleteAlamat)
}
