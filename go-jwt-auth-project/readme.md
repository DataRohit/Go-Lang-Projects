# Go JWT Auth Project

## Packages Used

1. github.com/gin-gonic/gin
2. github.com/joho/godotenv
3. github.com/go-playground/validator
4. go.mongodb.org/mongo-driver/mongo
5. golang.org/x/crypto/bcrypt
6. github.com/dgrijalva/jwt-go

## Routes

-   `/users/signup` - POST - handlers.SignUp to create a new user
-   `/users/login` - POST - handlers.Login to login the user and get the access and refresh jwt token

-   `/users` - GET - handlers.GetAllUsers to get all the users only accessible to ADMIN user type
-   `/uses/:userId` - GET - handlers.GetUserById to get a specific user only accessible to ADMIN user type

-   `/api/v1` - GET - Simple handler sending a success message accessible to all user types
-   `/api/v2` - GET - Simple handler sending a success message accessible to all user types

## Database

-   MongoDB (Docker Mongo Latest Image)

### Architecture Diagram

<img src="./architecture-diagram.svg" alt="Architecture Diagram" style="width:100%;"/>
