package router

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    movieRoutes "github.com/prathakpr/go_crud_app/internal/routes/movie"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api", logger.New())

        // Setup the Movie Routes
    movieRoutes.SetupMovieRoutes(api)
}
