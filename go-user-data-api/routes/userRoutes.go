package userRoutes

import (
	userController "go-user-data-api/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
}
