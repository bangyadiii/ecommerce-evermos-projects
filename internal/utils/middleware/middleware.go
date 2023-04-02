package middleware

import (
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/pkg/usecase"
	"ecommerce-evermos-projects/internal/utils/jwt"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// @TODO : make middleware like Auth

func VerifyToken(uc usecase.UsersUseCase, secret string) fiber.Handler {
	/*
		This code is used to authenticate a user and authorize them to access a certain endpoint.
		It first checks if the Authorization header contains "Bearer". If not, it returns an error response.
		It then uses authService to validate the tokenString, and if there is an error, it returns an error response.
		It then checks if the claims are valid and if not, it returns an error response.
		Finally, it uses userService to find the user by ID and check if the email matches with what is stored in payload.
		If all checks pass, it sets a current_user variable with the user object.
	*/

	return func(c *fiber.Ctx) error {
		token := c.Get("token")

		accessToken, err := jwt.ValidateToken(token, []byte(secret))
		if err != nil {
			return helper.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
		}

		u := accessToken.Claims.(jwt.CustomClaims)

		user, er := uc.GetUser(c.Context(), u.Email)

		if er != nil {
			return helper.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
		}

		if user.Email != u.Email {
			return helper.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
		}
		json, _ := json.Marshal(user)

		c.Set("current_user", string(json))
		return c.Next()
	}

}
