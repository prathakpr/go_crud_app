package movieRoutes

import "github.com/gofiber/fiber/v2"

//The function SetupMovieRoutes takes a fiber.Router and handles endpoints to the movie model
func SetupMovieRoutes(router fiber.Router) {
	movie := router.Group("/movie")
    // Create a Movie
    note.Post("/", func(c *fiber.Ctx) error {})
    // Read all Movie
    note.Get("/", func(c *fiber.Ctx) error {})
    // Read one Movie
    note.Get("/:movieId", func(c *fiber.Ctx) error {})
    // Update one Movie
    note.Put("/:movieId", func(c *fiber.Ctx) error {})
    // Delete one Movie
    note.Delete("/:movieId", func(c *fiber.Ctx) error {})
}