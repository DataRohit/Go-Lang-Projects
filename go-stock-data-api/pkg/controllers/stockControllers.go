package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/datarohit/go-stock-data-api/pkg/models"
	"github.com/datarohit/go-stock-data-api/pkg/schemas"
	"github.com/datarohit/go-stock-data-api/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	stocks, err := models.GetAllStocks()
	if err != nil {
		utils.LogError(r, "Unable to fetch stocks", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if len(stocks) == 0 {
		utils.LogError(r, "No stocks found", err)
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "no stocks found"})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "stocks fetched successfully",
		"stocks":  stocks,
	})
}

func GetStockBySymbol(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	vars := mux.Vars(r)
	symbol := vars["symbol"]

	stock, err := models.GetStockBySymbol(symbol)
	if err != nil {
		utils.LogError(r, "Unable to fetch stock", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "stock fetched successfully",
		"stock":   stock,
	})
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	var stock schemas.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		utils.LogError(r, "Error decoding request body", err)
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = utils.ValidateStock(stock)
	if err != nil {
		utils.LogError(r, "Validation failed", err)
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	createdStock, err := models.CreateStock(&stock)
	if err != nil {
		utils.LogError(r, "Error creating stock", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "stock created successfully",
		"stock":   createdStock,
	})
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	vars := mux.Vars(r)
	symbol := vars["symbol"]

	deletedStock, err := models.DeleteStock(symbol)
	if err != nil {
		utils.LogError(r, "Unable to delete stock", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "stock deleted successfully",
		"stock":   deletedStock,
	})
}
