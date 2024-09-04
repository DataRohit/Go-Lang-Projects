package routes

import (
	"go-simple-crm-tool/api/handlers"

	"github.com/gorilla/mux"
)

var RegisterLeadRoutes = func(router *mux.Router) {
	router.HandleFunc("/leads", handlers.CreateLeadsHandler).Methods("POST")
	router.HandleFunc("/leads", handlers.GetAllLeadsHandler).Methods("GET")
	// router.HandleFunc("/leads/{id}", handlers.GetLeadByIDHandler).Methods("GET")
	// router.HandleFunc("/leads/{id}", handlers.DeleteLeadHandler).Methods("DELETE")
	// router.HandleFunc("/leads/{id}", handlers.UpdateLeadHandler).Methods("PUT")
}
