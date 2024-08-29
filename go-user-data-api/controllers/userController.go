package userController

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	userModels "github.com/datarohit/go-user-data-api/models"
	"github.com/datarohit/go-user-data-api/utils"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	users, err := userModels.GetAllUsers(ctx)
	if err != nil {
		utils.LogError(r, "Unable to fetch users", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error fetching users"})
		return
	}

	res, err := json.Marshal(users)
	if err != nil {
		utils.LogError(r, "Error marshalling response", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error processing response"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
