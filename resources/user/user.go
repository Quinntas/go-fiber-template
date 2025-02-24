package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/quinntas/go-fiber-template/database"
	"github.com/quinntas/go-fiber-template/database/repository"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(repository.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "failed to parse body",
		})
	}
	createUser, err := database.Repo.CreateUser(context.Background(), repository.CreateUserParams{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(createUser)
}
