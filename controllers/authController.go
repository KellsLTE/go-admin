package controllers

import (
	"github.com/KellsLTE/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "Richards",
		LastName: "Collider",
		Email: "colliderrichards@gmail.com",
	}

	return c.JSON(user)
}