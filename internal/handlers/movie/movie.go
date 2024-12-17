package movieHandler

//Adding the Read Movies handler -
func GetMovies(c *fiber.Ctx) error {
    db := database.DB
    var movie []model.movie

    // find all movies in the database
    db.Find(&movies)

    // If no movie is present return an error
    if len(movies) == 0 {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No movies present", "data": nil})
    }

    // Else return movies
    return c.JSON(fiber.Map{"status": "success", "message": "Movies Found", "data": movies})
}

//Adding the Create Movies handler -
func CreateMovies(c *fiber.Ctx) error {
    db := database.DB
    movie := new(model.Movie)

    // Store the body in the movie and return error if encountered
    err := c.BodyParser(movie)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }
    // Add a uuid to the movie
    movie.ID = uuid.New()
    // Create the Movie and return error if encountered
    err = db.Create(&movie).Error
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create movie", "data": err})
    }

    // Return the created movie
    return c.JSON(fiber.Map{"status": "success", "message": "Created Movie", "data": movie})
}

//Adding the Get Specific Movie Handler
func GetMovie(c *fiber.Ctx) error {
    db := database.DB
    var movie model.Movie

    // Read the param movieId
    id := c.Params("movieId")

    // Find the movie with the given Id
    db.Find(&movie, "id = ?", id)

    // If no such movie present return an error
    if movie.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No movie present", "data": nil})
    }

    // Return the movie with the Id
    return c.JSON(fiber.Map{"status": "success", "message": "Movies Found", "data": movie})
}

//Adding the Update Movie Handler

func UpdateMovie(c *fiber.Ctx) error {
    type updateMovie struct {
        Title    string `json:"title"`
        SubTitle string `json:"sub_title"`
        Text     string `json:"Text"`
    }
    db := database.DB
    var movie model.Movie

    // Read the param movieId
    id := c.Params("movieId")

    // Find the movie with the given Id
    db.Find(&movie, "id = ?", id)

    // If no such movie present return an error
    if movie.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No movie present", "data": nil})
    }

    // Store the body containing the updated data and return error if encountered
    var updateMovieData updateMovie
    err := c.BodyParser(&updateMovieData)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }

    // Edit the movie
    movie.Title = updateMovieData.Title
    movie.SubTitle = updateMovieData.SubTitle
    movie.Text = updateMovieData.Text

    // Save the Changes
    db.Save(&movie)

    // Return the updated movie
    return c.JSON(fiber.Map{"status": "success", "message": "Movies Found", "data": movie})
}

//Adding the Delete Movie Handler
func DeleteMovie(c *fiber.Ctx) error {
    db := database.DB
    var movie model.Movie

    // Read the param movieId
    id := c.Params("movieId")

    // Find the movie with the given Id
    db.Find(&movie, "id = ?", id)

    // If no such movie present return an error
    if movie.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No movie present", "data": nil})
    }

    // Delete the movie and return error if encountered
    err := db.Delete(&movie, "id = ?", id).Error

    if err != nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete movie", "data": nil})
    }

    // Return success message
    return c.JSON(fiber.Map{"status": "success", "message": "Deleted Movie"})
}
