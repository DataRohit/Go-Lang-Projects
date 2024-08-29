package userController

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	userModels "github.com/datarohit/go-user-data-api/models"
	"github.com/datarohit/go-user-data-api/schemas"
	"github.com/datarohit/go-user-data-api/utils"
	"github.com/gorilla/mux"
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

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	vars := mux.Vars(r)
	id := vars["id"]

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, err := userModels.GetUserByID(ctx, id)
	if err != nil {
		utils.LogError(r, "Unable to fetch user", err)
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		utils.LogError(r, "Error marshalling response", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error processing response"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	var user schemas.User
	if err := utils.ParseBody(r, &user); err != nil {
		utils.LogError(r, "Error parsing request body", err)
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	createdUser, err := userModels.CreateUser(ctx, user)
	if err != nil {
		utils.LogError(r, "Error creating user", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error creating user"})
		return
	}

	res, err := json.Marshal(createdUser)
	if err != nil {
		utils.LogError(r, "Error marshalling response", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error processing response"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	vars := mux.Vars(r)
	id := vars["id"]

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	deletedUser, err := userModels.DeleteUser(ctx, id)
	if err != nil {
		utils.LogError(r, "Error deleting user", err)
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	res, err := json.Marshal(deletedUser)
	if err != nil {
		utils.LogError(r, "Error marshalling response", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error processing response"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
