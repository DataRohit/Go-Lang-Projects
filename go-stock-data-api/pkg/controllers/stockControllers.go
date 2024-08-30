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
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error fetching stocks"})
		return
	}

	res, err := json.Marshal(stocks)
	if err != nil {
		utils.LogError(r, "Error marshalling response", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error processing response"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStockBySymbol(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	vars := mux.Vars(r)
	symbol := vars["symbol"]

	stock, err := models.GetStockBySymbol(symbol)
	if err != nil {
		utils.LogError(r, "Unable to fetch stock", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error fetching stock"})
		return
	}

	if stock == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Stock not found"})
		return
	}

	res, err := json.Marshal(stock)
	if err != nil {
		utils.LogError(r, "Error marshalling response", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error processing response"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	var stock schemas.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		utils.LogError(r, "Error decoding request body", err)
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
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
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error creating stock"})
		return
	}

	res, err := json.Marshal(createdStock)
	if err != nil {
		utils.LogError(r, "Error marshalling response", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error processing response"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
