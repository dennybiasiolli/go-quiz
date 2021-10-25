package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/dennybiasiolli/go-quiz/auth"
	"github.com/dennybiasiolli/go-quiz/common"
	"github.com/dennybiasiolli/go-quiz/quiz"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	common.GetEnvVariables(".env", ".env.default")

	common.ConnectDb()
	common.GetDB().AutoMigrate(
		&auth.User{},
		&quiz.Quiz{},
		&quiz.Question{},
	)

	app := fiber.New()
	app.Use(cors.New())

	setupFiberRoutes(app)

	app.Get("/", auth.GetJwtAuthMiddleware(), func(c *fiber.Ctx) error {
		return c.JSON(c.Locals("user"))
	})

	log.Fatal(app.Listen(common.HTTP_LISTEN))
}
