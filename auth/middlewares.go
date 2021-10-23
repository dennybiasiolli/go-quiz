package auth

import (
	"github.com/dennybiasiolli/go-quiz/common"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func GetJwtAuthMiddleware() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(common.JWT_HMAC_SAMPLE_SECRET),
		SuccessHandler: func(c *fiber.Ctx) error {
			u := c.Locals("user").(*jwt.Token)
			claims := u.Claims.(jwt.MapClaims)

			if claims["token_type"] != "access" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token type"})
			}

			db := common.GetDB()
			var user User
			err := db.Where(User{IsActive: true}).First(&user, claims["user_id"]).Error
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			}
			c.Locals("user", user)
			return c.Next()
		},
	})
}
