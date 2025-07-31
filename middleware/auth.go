package middleware

import (
	"fmt"
	"strings"
	"task-manager/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{"error": "Missing or malformed JWT"})
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(config.JWT_SECRET), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{"error": "Invalid or expired JWT"})
		}
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(float64)
		c.Locals("user_id", uint(userID))
		return c.Next()
	}
}
