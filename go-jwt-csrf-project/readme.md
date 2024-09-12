# Go JWT Auth Project

## Packages Used

1. github.com/dgrijalva/jwt-go
2. github.com/justinas/alice
3. golang.org/x/crypto

## Routes

-   `/restricted` - GET - Renders the restricted template and returns the csrf secret token
-   `/login` - GET - Renders the login templates
-   `/login` - POST - Handles the form submit to log the user and set the auth cookies
-   `/register` - GET - Renders the register templates
-   `/register` - POST - Handles the form submit to register the user and set the auth cookies
-   `/logout` - GET - Remove the auth cookies and redirect to login page
-   `/deleteUser` - GET - Remove the auth cookies and delete the user

### Architecture Diagram

<img src="./architecture-diagram.svg" alt="Architecture Diagram" style="width:100%;"/>
