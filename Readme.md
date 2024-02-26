# Movie CRUD API

This is a simple CRUD (Create, Read, Update, Delete) API for movies, written in Go.

## Getting Started

To run this project, you will need to have Go installed on your machine. You can download it from the [official website](https://golang.org/dl/).

## Running the Project

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Run `go run main.go` to start the server.

The server will start on port 8080.

## API Endpoints

The API has the following endpoints:

- `GET /movies`: Returns a list of all movies.
- `GET /movie/{id}`: Returns the movie with the specified ID.
- `POST /movies`: Adds a new movie. The movie details should be included in the request body in JSON format.
- `PUT /movies/{id}`: Updates the movie with the specified ID. The updated movie details should be included in the request body in JSON format.
- `DELETE /movies/{id}`: Deletes the movie with the specified ID.

## Data Model

The API uses the following data model:

```go
type Movie struct {
    Id       string `json:"id"`       // Unique identifier of the movie.
    Name     string `json:"name"`     // Name of the movie.
    Title     string `json:"title"`     // Year of release of the movie.
    Director *Director `json:"director"` // Director of the movie.
}

type Director struct {
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
}



## Dependencies

This project uses the `gorilla/mux` package for routing. You can install it with `go get -u github.com/gorilla/mux`.
