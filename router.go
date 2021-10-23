package main

import (
	"github.com/dennybiasiolli/go-quiz/auth"
	"github.com/gofiber/fiber/v2"
)

func setupFiberRoutes(app *fiber.App) {
	auth.JwtTokenRegister(app.Group("/token"))
	auth.Oauth2GoogleRegister(app.Group("/oauth2"))
}
