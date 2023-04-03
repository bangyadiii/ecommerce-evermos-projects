package http

import (
	route "ecommerce-evermos-projects/internal/server/http/handler"

	"ecommerce-evermos-projects/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/api/v1") // /api

	route.AuthRoute(api, containerConf)
	route.BooksRoute(api, containerConf)
	route.CategoryRoute(api, containerConf)
	route.UserRoute(api, containerConf)
}
