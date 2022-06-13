package controllers

import(
	"strconv"

	"github.com/KellsLTE/go-admin/database"
	"github.com/KellsLTE/go-admin/models"
	"github.com/gofiber/fiber/v2"
	//"golang.org/x/crypto/bcrypt"
)

// function that pulls all the roles from the datanase
func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role // SLice that will contain the roles

	database.DB.Find(&roles)

	return c.JSON(roles)
}

//function that creates a role
func CreateRole(c *fiber.Ctx) error {
	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, p := range list {
		id, _ := strconv.Atoi(p.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name: roleDTO["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

//function to get single role
func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

// function to update role record
func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, p := range list {
		id, _ := strconv.Atoi(p.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	//pull details from the database
	result := models.Role{
		Name: roleDTO["name"].(string),
		Permissions: permissions,
	}
	
	//delete the details from the database
	database.DB.Table("role_permissions").Where("role_id = ?", id).Delete(&result)

	//pull the details to be updated from the database
	role := models.Role{
		Id: uint(id),
		Name: roleDTO["name"].(string),
		Permissions: permissions,
	}

	//update the details
	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

//function to delete role record
func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return nil
}