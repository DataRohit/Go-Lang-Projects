# Go Movies CRUD API

## Packages Used

1. gorilla/mux
2. gorm.io/gorm
3. gorm.io/driver/sqlite
4. google/uuid
5. net/http
6. log
7. encoding/json

## Routes

-   `/books` - GetBooks() handler to fetch all the available books
-   `/books/{id}` - GetBookByID() handler to fetch a single book using the `bookId`
-   `/books` - CreateBook() handler to create a new book with unique uuid
-   `/books/{id}` - DeleteBook() handler to delete an instance of book from the db
-   `/books` - UpdateBook() handler to update an existing book record in db

## Database

-   Sqlite DB

### Folder Structure

<img src="./folder-structure.svg" alt="Folder Structure" style="width:100%;"/>

### Architecture Diagram

<img src="./architecture-diagram.svg" alt="Architecture Diagram" style="width:100%;"/>
