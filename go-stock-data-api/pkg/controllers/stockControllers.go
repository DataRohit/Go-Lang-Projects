package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/datarohit/go-stock-data-api/pkg/models"
	"github.com/datarohit/go-stock-data-api/pkg/utils"
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
