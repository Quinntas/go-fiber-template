package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "go-fiber-template",
		CaseSensitive: false,
		ServerHeader:  "go-fiber-template",
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(cors.New())
	app.Use(helmet.New())

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "Too many requests",
			})
		},
	}))

	app.Use(logger.New())

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Use(recover.New())

	app.Use(requestid.New())

	v1Router := app.Group("/api/v1")

	v1Router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "ok",
		})
	})

	userRouter := v1Router.Group("/user")

	userRouter.Use(func(c *fiber.Ctx) error {
		c.Locals("user", "admin")
		return c.Next()
	})

	userRouter.Get("/", func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(string)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
		fmt.Println(user)
		return c.JSON(fiber.Map{
			"message": "ok",
		})
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
