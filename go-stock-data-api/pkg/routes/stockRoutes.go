package routes

import (
	"github.com/datarohit/go-stock-data-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterStockRoutes = func(router *mux.Router) {
	router.HandleFunc("/stocks", controllers.GetAllStocks).Methods("GET", "OPTIONS")
}
