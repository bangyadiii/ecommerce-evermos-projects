package handler

import (
	"ecommerce-evermos-projects/internal/infrastructure/container"
	"ecommerce-evermos-projects/internal/pkg/controller"
	"ecommerce-evermos-projects/internal/pkg/repository"
	"ecommerce-evermos-projects/internal/pkg/usecase"
	"ecommerce-evermos-projects/internal/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewCategoryRepository(containerConf.Mysqldb)
	userRepo := repository.NewUserRepository(containerConf.Mysqldb)
	uc := usecase.NewCategoryUseCase(repo)
	ucUser := usecase.NewUsersUseCase(userRepo, *containerConf.Apps)

	controller := controller.NewCategoryController(uc)
	secret := containerConf.Apps.SecretJwt
	categoryAPI := r.Group("/category")
	categoryAPI.Get("/", controller.GetAllCategory)
	categoryAPI.Get("/:category_id", controller.GetCategoryByID)
	categoryAPI.Post("/", middleware.VerifyToken(ucUser, secret), controller.CreateCategory)
	categoryAPI.Put("/:category_id", middleware.VerifyToken(ucUser, secret), controller.UpdateCategoryByID)
	categoryAPI.Delete("/:category_id", middleware.VerifyToken(ucUser, secret), controller.DeleteCategoryByID)
}
