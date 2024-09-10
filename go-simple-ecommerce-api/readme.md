# Go JWT Auth Project

## Packages Used

1. github.com/go-playground/validator/v10
2. github.com/golang-jwt/jwt/v5
3. github.com/golang-migrate/migrate/v4
4. github.com/gorilla/mux
5. github.com/joho/godotenv
6. github.com/go-sql-driver/mysql
7. golang.org/x/crypt

## Routes

-   `/login` - POST - h.handleLogin function login an existing user using the email and password
-   `/register` - POST - h.handleRegister function to register a new user with unique email address
-   `/users/{userID}` - GET - JWTAuth protected h.handleGetUser function to get the specific user details using the userID

-   `/products` - GET - h.handleGetProducts function to get all the products available in in the inventory
-   `/products/{productID}` - GET - h.handleGetProduct function to get details of specific product using the product id
-   `/products` - POST - JWTAuth protected h.handleCreateProduct function to crete a new product

-   `/cart/checkout` - POST - JWTAuth protected h.handleCheckout function to place a new order for the authenticated user

## Database

-   MySQL (Docker MySQL Latest Image)

### Architecture Diagram

<img src="./architecture-diagram.svg" alt="Architecture Diagram" style="width:100%;"/>
