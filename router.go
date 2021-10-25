package main

import (
	"github.com/dennybiasiolli/go-quiz/auth"
	"github.com/dennybiasiolli/go-quiz/quiz"
	"github.com/gofiber/fiber/v2"
)

func setupFiberRoutes(app *fiber.App) {
	auth.JwtTokenRegister(app.Group("/token"))
	auth.Oauth2GoogleRegister(app.Group("/oauth2"))
	quiz.QuizRoutesRegister(app.Group("/quiz").Use(auth.GetJwtAuthMiddleware()))
}
