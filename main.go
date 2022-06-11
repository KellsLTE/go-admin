package main

import (
	"log"

	"github.com/KellsLTE/go-admin/database"
	"github.com/KellsLTE/go-admin/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

    app := fiber.New()

    routes.Setup(app)

    log.Fatal(app.Listen(":8000"))
}