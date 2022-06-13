package middleware

import (
	"github.com/KellsLTE/go-admin/utils"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt-token")

	if _, err := utils.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized user",
		})
	}

	return c.Next()
}