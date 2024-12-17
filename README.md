# Go_Crud_App: Building RESTful APIs on Golang (Feat. GORM)

## Project Overview

This project is a RESTful API application built using **Go** and the **Fiber** framework, with **GORM** for database interaction. It demonstrates how to create and manage API endpoints that interact with a **PostgreSQL** database. The project focuses on implementing basic CRUD operations for a "movies" resource, using **PostgreSQL** for database storage.

Key Features:
- **CRUD Operations for Movies**: Create, read, update, and delete movies.
- **Go Fiber Framework**: Lightweight and fast web framework for building APIs.
- **GORM**: ORM for interacting with PostgreSQL.
- **Database Integration**: Simple database setup with PostgreSQL.

### Key Features

#### Movies Module
- **Manage Movies**: Store and manage movie details such as title, director, release date, genres, and rating.
- **Perform CRUD Operations on Movies**:
  - `POST /movies`: Create a new movie.
  - `GET /movies`: Fetch all movies.
  - `GET /movies/:id`: Fetch a movie by its ID.
  - `PUT /movies/:id`: Update an existing movie.
  - `DELETE /movies/:id`: Delete a movie.
- Uses **PostgreSQL** for storing movie data with relational integrity.

### Technologies Used:
- **Go** for backend development.
- **Fiber** as the web framework for handling HTTP requests.
- **GORM** for object-relational mapping (ORM).
- **PostgreSQL** for database storage.

## Setup Instructions

### Clone the Repository
To get started with this project, clone the repository to your local machine:

```bash
git clone https://github.com/prathakpr/go_crud_app.git
cd go-fiber-api
