package userRoutes

import (
	userController "github.com/datarohit/go-user-data-api/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
}
