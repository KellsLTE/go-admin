package controllers

import (
	"strconv"

	"github.com/KellsLTE/go-admin/database"
	"github.com/KellsLTE/go-admin/models"
	"github.com/gofiber/fiber/v2"
	//"golang.org/x/crypto/bcrypt"
)

// function that pulls all the users from the datanase
func AllUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	
	limit := 5
	offset := (page - 1) * limit
	var total int64

	var users []models.User // SLice that will contain the users

	database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)

	database.DB.Model(&models.User{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total": total,
			"page": page,
			"last_page": float64(int(total)/ limit),
		},
	})
}

//function that creates a user
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("password")

	database.DB.Create(&user)

	return c.JSON(user)
}

//function to get single user
func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Find(&user)

	return c.JSON(user)
}

// function to update user record
func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

//function to delete user record
func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}