# Go Users Data API

## Packages Used

1. gorilla/mux
2. net/http
3. os/signal
4. context
5. log
6. encoding/json
7. go.mongodb.org/mongo-driver

## Routes

-   `/users` - GetUsers() handler to fetch all the available users
-   `/users/{id}` - GetUserByID() handler to fetch a single user using the `id`
-   `/users` - CreateUser() handler to create a new user with unique uuid
-   `/users/{id}` - DeleteUser() handler to delete an instance of user from the db
-   `/users/{id}` - UpdateUser() handler to update an existing user record in db

## Database

-   MongoDB

### Folder Structure

<img src="./folder-structure.svg" alt="Folder Structure" style="width:100%;"/>

### Architecture Diagram

<img src="./architecture-diagram.svg" alt="Architecture Diagram" style="width:100%;"/>
