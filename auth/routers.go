package auth

import "github.com/gofiber/fiber/v2"

func JwtTokenRegister(router fiber.Router) {
	router.Post("/refresh/", TokenRefresh)
}

func Oauth2GoogleRegister(router fiber.Router) {
	router.Get("/google/login", GoogleOauth2Login)
	router.Get("/google/callback", GoogleOauth2Callback)
	router.Get("/google", GoogleOauth2)
}
