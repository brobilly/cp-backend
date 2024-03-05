package middleware

import (
	"fmt"
	"strings"

	"campus-api/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticate(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	if authorizationHeader != "" {
		token := strings.TrimPrefix(authorizationHeader, "Bearer ")
		fmt.Println("Token from Authorization Header:", token)
		if _, err := util.ParseJwt(token); err == nil {
			return c.Next()
		}
	}

	return c.Next()
}
