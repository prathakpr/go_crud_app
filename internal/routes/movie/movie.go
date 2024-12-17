package movieRoutes

import(
	"github.com/gofiber/fiber/v2"
	movieHandler "github.com/prathakpr/go_crud_app/internal/handlers/movie"
)

//The function SetupMovieRoutes takes a fiber.Router and handles endpoints to the movie model
func SetupMovieRoutes(router fiber.Router) {
	movie := router.Group("/movie")
    // Create a Movie
    movie.Post("/", movieHandler.CreateMovies)
    // Read all Movie
    movie.Get("/", movieHandler.GetMovies)
    // Read one Movie
    movie.Get("/:movieId", movieHandler.GetMovie)
    // Update one Movie
    movie.Put("/:movieId", movieHandler.UpdateMovie)
    // Delete one Movie
    movie.Delete("/:movieId", movieHandler.DeleteMovie)
}