package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UserSerializer(user User) fiber.Map {
	return fiber.Map{
		"id":        user.ID,
		"full_name": strings.TrimSpace(user.FirstName + " " + user.LastName),
	}
}
