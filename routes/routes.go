package routes

import (
	"github.com/KellsLTE/go-admin/controllers"
	"github.com/KellsLTE/go-admin/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("api/register", controllers.Register)

	app.Post("api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)

	// Authentication Routes
	app.Put("/api/users/profile", controllers.UpdateProfileData)
	app.Put("/api/users/password", controllers.UpdatePassword)
	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)

	//User Routes
	app.Get("api/users", controllers.AllUsers)
	app.Post("api/users", controllers.CreateUser)
	app.Get("api/users/:id", controllers.GetUser)
	app.Put("api/users/:id", controllers.UpdateUser)
	app.Delete("api/users/:id", controllers.DeleteUser)

	//Role Routes
	app.Get("api/roles", controllers.AllRoles)
	app.Post("api/roles", controllers.CreateRole)
	app.Get("api/roles/:id", controllers.GetRole)
	app.Put("api/roles/:id", controllers.UpdateRole)
	app.Delete("api/roles/:id", controllers.DeleteRole)

	app.Get("api/permissions", controllers.AllPermissions)

	//Product Routes
	app.Get("api/products", controllers.AllProducts)
	app.Post("api/products", controllers.CreateProduct)
	app.Get("api/products/:id", controllers.GetProduct)
	app.Put("api/products/:id", controllers.UpdateProduct)
	app.Delete("api/products/:id", controllers.DeleteProduct)

	//File upload and donwload
	app.Post("api/upload", controllers.Upload)
	app.Static("api/uploads", "./storage/uploads")

	//Order	Routes
	app.Get("/api/orders", controllers.AllOrders)
	app.Post("/api/export", controllers.Export)
}