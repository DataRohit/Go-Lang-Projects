# Go Movies CRUD API

# Go Movies CRUD API

## Packages Used

1. gorilla/mux
2. google/uuid
3. net/http
4. log
5. encoding/json

## Routes

-   `/movies` - getMovies() handler to fetch all the available movies
-   `/movies/{id}` - getMovie() handler to fetch a single movie using the `id`
-   `/movies` - createMovie() handler to create a new movie with unique uuid
-   `/movies/{id}` - deleteMovie() handler to delete an instance of movie from the slice
-   `/movies/{id}` - updateMovie() handler to update an existing book record in slice

### Architecture Diagram

<img src="./architecture-diagram.svg" alt="Architecture Diagram" style="width:100%;"/>
