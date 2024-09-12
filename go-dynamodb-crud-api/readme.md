# Go JWT Auth Project

## Packages Used

1. github.com/aws/aws-sdk-go
2. github.com/go-chi/chi
3. github.com/go-chi/chi/v5
4. github.com/go-chi/cors
5. github.com/go-ozzo/ozzo-validation/v4
6. github.com/stretchr/testify
7. go.uber.org/zap

## Routes

-   `/health` - GET - HealthHandler.Get to get the status if the database is alive or not
-   `/health` - POST | PUT | DELETE - HealthHander.Post, HealthHandler.Put & HealthHandler.Delete to handle unsupported methods
-   `/health` - OPTIONS - HealthHandler.Options to handle the options requiest to the endpoint

<br />

-   `/product` - POST - ProductHandler.Post to handle creating a new product
-   `/product` - GET - ProductHandler.Get to handle retrieving all products
-   `/product/{ID}` - GET - ProductHandler.Get to handle retrieving a single product by ID
-   `/product/{ID}` - PUT - ProductHandler.Put to handle updating a single product by ID
-   `/product/{ID}` - DELETE - ProductHandler.Delete to handle deleting a single product by ID
-   `/product` - OPTIONS - ProductHandler.Options to handle the Options request to the endpoint

## Database

-   DynamoDB (Amazon DynamoDB Latest Image)

### Architecture Diagram

<img src="./architecture-diagram.svg" alt="Architecture Diagram" style="width:100%;"/>
