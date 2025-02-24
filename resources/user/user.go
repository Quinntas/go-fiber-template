package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/quinntas/go-fiber-template/database"
	"gorm.io/gorm"
)

type User struct {
	database.DefaultModel
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []User
	db.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user User
	id := c.Params("id")
	err := db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "record not found",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.JSON(user)
}

func NewUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse body")
	}
	db.Create(&user)
	return c.JSON(user)
}
