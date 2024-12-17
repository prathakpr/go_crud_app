package movieHandler

//Adding the Read Movies handler -
func GetMovies(c *fiber.Ctx) error {
    db := database.DB
    var movie []model.movie

    // find all movies in the database
    db.Find(&movies)

    // If no note is present return an error
    if len(movies) == 0 {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No movies present", "data": nil})
    }

    // Else return notes
    return c.JSON(fiber.Map{"status": "success", "message": "Movies Found", "data": movies})
}

//Adding the Create Movies handler -
func CreateMovies(c *fiber.Ctx) error {
    db := database.DB
    note := new(model.Movie)

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
    return c.JSON(fiber.Map{"status": "success", "message": "Created Movie", "data": note})
}

//Adding the Get Specific Movie Handler
func GetMovie(c *fiber.Ctx) error {
    db := database.DB
    var note model.Movie

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

