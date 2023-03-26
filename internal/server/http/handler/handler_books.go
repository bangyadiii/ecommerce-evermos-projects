package handler

import (
	"ecommerce-evermos-projects/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	bookscontroller "ecommerce-evermos-projects/internal/pkg/controller"

	booksrepository "ecommerce-evermos-projects/internal/pkg/repository"

	booksusecase "ecommerce-evermos-projects/internal/pkg/usecase"
)

func BooksRoute(r fiber.Router, containerConf *container.Container) {
	repo := booksrepository.NewBooksRepository(containerConf.Mysqldb)
	usecase := booksusecase.NewBooksUseCase(repo)
	controller := bookscontroller.NewBooksController(usecase)

	booksAPI := r.Group("/books")
	booksAPI.Get("", controller.GetAllBooks)
	booksAPI.Get("/:id_books", controller.GetBooksByID)
	booksAPI.Post("", controller.CreateBooks)
	booksAPI.Put("/:id_books", controller.UpdateBooksByID)
	booksAPI.Delete("/:id_books", controller.DeleteBooksByID)
}
