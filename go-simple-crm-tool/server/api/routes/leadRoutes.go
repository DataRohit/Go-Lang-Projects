package routes

import (
	"go-simple-crm-tool/api/handlers"

	"github.com/gorilla/mux"
)

var RegisterLeadRoutes = func(router *mux.Router) {
	router.HandleFunc("/leads", handlers.CreateLeadsHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/leads", handlers.GetAllLeadsHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/leads", handlers.DeleteLeadsHandler).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/leads", handlers.UpdateLeadsHandler).Methods("PUT", "OPTIONS")
}
