package routes

import (
	"github.com/datarohit/go-stock-data-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterStockRoutes = func(router *mux.Router) {
	router.HandleFunc("/stocks", controllers.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/stocks/{symbol}", controllers.GetStockBySymbol).Methods("GET", "OPTIONS")
	router.HandleFunc("/stocks", controllers.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/stocks/{symbol}", controllers.DeleteStock).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/stocks/{symbol}", controllers.UpdateStock).Methods("PUT", "OPTIONS")
}
