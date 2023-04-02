package handler

import (
	"ecommerce-evermos-projects/internal/infrastructure/container"
	"ecommerce-evermos-projects/internal/pkg/controller"
	 "ecommerce-evermos-projects/internal/pkg/repository"
	 "ecommerce-evermos-projects/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewUserRepository(containerConf.Mysqldb)
	usecase := usecase.NewUsersUseCase(repo, *containerConf.Apps)
	controller := controller.NewAuthControllerImpl(usecase)

	booksAPI := r.Group("/auth")
	booksAPI.Post("/login", controller.Login)
	booksAPI.Post("/register", controller.Register)
}
