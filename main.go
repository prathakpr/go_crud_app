package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/prathakpr/go_crud_app/database"
	"github.com/prathakpr/go_crud_app/router"
)

func main() {
    // Start a new fiber app
    app := fiber.New()

	// Connect to the Database
    database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

    // Listen on PORT 300
    app.Listen(":3000")
}
