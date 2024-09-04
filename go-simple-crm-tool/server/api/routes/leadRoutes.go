package routes

import (
	"go-simple-crm-tool/api/handlers"

	"github.com/gorilla/mux"
)

var RegisterLeadRoutes = func(router *mux.Router) {
	router.HandleFunc("/leads", handlers.CreateMultipleLeadsHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/leads", handlers.GetAllLeadsHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/leads", handlers.DeleteMultipleLeadsHandler).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/leads", handlers.UpdateMultipleLeadsHandler).Methods("PUT", "OPTIONS")

	router.HandleFunc("/lead", handlers.CreateSingleLeadHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/lead", handlers.GetRandomLeadHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/lead/{id}", handlers.GetLeadByIDHandler).Methods("GET", "OPTIONS")
}
