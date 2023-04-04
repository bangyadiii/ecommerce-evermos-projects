package handler

import (
	"ecommerce-evermos-projects/internal/infrastructure/container"
	"ecommerce-evermos-projects/internal/pkg/controller"
	"ecommerce-evermos-projects/internal/pkg/repository"
	"ecommerce-evermos-projects/internal/pkg/usecase"
	"ecommerce-evermos-projects/internal/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

func TokoRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewTokoRepository(containerConf.Mysqldb)
	userRepo := repository.NewUserRepository(containerConf.Mysqldb)
	uc := usecase.NewTokoUseCase(repo)
	ucUser := usecase.NewUsersUseCase(userRepo, *containerConf.Apps)

	controller := controller.NewTokoController(uc)
	secret := containerConf.Apps.SecretJwt
	tokoAPI := r.Group("/toko")

	tokoAPI.Get("/my", middleware.VerifyToken(ucUser, secret), controller.GetMyToko)
	tokoAPI.Get("/:toko_id", middleware.VerifyToken(ucUser, secret), controller.GetTokoByID)
	tokoAPI.Put("/:toko_id", middleware.VerifyToken(ucUser, secret), controller.UpdateToko)
	tokoAPI.Get("/", controller.FetchToko)
}
