package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quinntas/go-fiber-template/modules/shared/resources/healthCheck"
)

func InitRouter(group fiber.Router) {
	group.Get("/", healthCheck.Controller)
}
