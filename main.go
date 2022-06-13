package main

import (
	"fmt"
	"log"

	"github.com/KellsLTE/go-admin/config"
	"github.com/KellsLTE/go-admin/database"
	"github.com/KellsLTE/go-admin/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

    routes.Setup(app)

	env := config.Env("APP_ENVIRONMENT")
	port := config.Env("APP_PORT")

	fmt.Println("Server is running in " + env + " mode on port " + port)

    log.Fatal(app.Listen(":" + port))
}