package middleware

import (
	"errors"
	"strconv"

	"github.com/KellsLTE/go-admin/database"
	"github.com/KellsLTE/go-admin/models"
	"github.com/KellsLTE/go-admin/utils"
	"github.com/gofiber/fiber/v2"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt-token")

	id, err := utils.ParseJwt(cookie)

	if err != nil{
		return err
	}

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	database.DB.Preload("Permissions").Find(&role)

	if c.Method() == "GET" {
		for _, permission := range role.Permissions{
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions{
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("Unauthorized")
}