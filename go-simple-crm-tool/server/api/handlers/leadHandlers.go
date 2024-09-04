package handlers

import (
	"encoding/json"
	"fmt"
	"go-simple-crm-tool/api/models"
	"go-simple-crm-tool/api/schemas"
	"go-simple-crm-tool/pkg/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func CreateMultipleLeadsHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info(logrus.Fields{
		"action":     "CreateMultipleLeadsHandler",
		"method":     r.Method,
		"url":        r.URL.Path,
		"remoteAddr": r.RemoteAddr,
	}, "POST request received at /leads")

	var leads []schemas.Lead

	err := utils.ParseBody(r, &leads)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "CreateMultipleLeadsHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Error decoding request body: %v", err))
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	var createdLeads []schemas.Lead
	for _, lead := range leads {
		err = utils.ValidateLead(&lead)
		if err != nil {
			utils.Warn(logrus.Fields{
				"action":     "CreateMultipleLeadsHandler",
				"method":     r.Method,
				"url":        r.URL.Path,
				"remoteAddr": r.RemoteAddr,
			}, fmt.Sprintf("Validation failed: %v", err))
			utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		createdLead, err := models.CreateLead(&lead)
		if err != nil {
			utils.Warn(logrus.Fields{
				"action":     "CreateMultipleLeadsHandler",
				"method":     r.Method,
				"url":        r.URL.Path,
				"remoteAddr": r.RemoteAddr,
			}, fmt.Sprintf("Error creating lead: %v", err))
			utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		createdLeads = append(createdLeads, *createdLead)
	}

	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "leads created successfully",
		"leads":   createdLeads,
	})
}

func GetAllLeadsHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info(logrus.Fields{
		"action":     "GetAllLeadsHandler",
		"method":     r.Method,
		"url":        r.URL.Path,
		"remoteAddr": r.RemoteAddr,
	}, "GET request received at /leads")

	leads, err := models.GetAllLeads()
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "GetAllLeadsHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Unable to fetch leads: %v", err))
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if len(leads) == 0 {
		utils.Warn(logrus.Fields{
			"action":     "GetAllLeadsHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("No leads found: %v", err))
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "no leads found"})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "leads fetched successfully",
		"leads":   leads,
	})
}

func DeleteMultipleLeadsHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info(logrus.Fields{
		"action":     "DeleteMultipleLeadsHandler",
		"method":     r.Method,
		"url":        r.URL.Path,
		"remoteAddr": r.RemoteAddr,
	}, "DELETE request received at /leads")

	var ids []uuid.UUID

	err := utils.ParseBody(r, &ids)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "DeleteMultipleLeadsHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Error decoding request body: %v", err))
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	for _, id := range ids {
		err = models.DeleteLead(id)
		if err != nil {
			utils.Warn(logrus.Fields{
				"action":     "DeleteMultipleLeadsHandler",
				"method":     r.Method,
				"url":        r.URL.Path,
				"remoteAddr": r.RemoteAddr,
				"leadID":     id,
			}, fmt.Sprintf("Error deleting lead with id %s: %v", id, err))
			utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("error deleting lead with id %s: %v", id, err),
			})
			return
		}
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{
		"message": "leads deleted successfully",
	})
}

func UpdateMultipleLeadsHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info(logrus.Fields{
		"action":     "UpdateMultipleLeadsHandler",
		"method":     r.Method,
		"url":        r.URL.Path,
		"remoteAddr": r.RemoteAddr,
	}, "PUT request received at /leads")

	var updateRequests []schemas.UpdateLeadRequest

	err := utils.ParseBody(r, &updateRequests)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "UpdateMultipleLeadsHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Error decoding request body: %v", err))
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	var updatedLeads []schemas.Lead
	for _, req := range updateRequests {
		updatedLead, err := models.UpdateLeadByID(req.ID, req.Data)
		if err != nil {
			utils.Warn(logrus.Fields{
				"action":     "UpdateMultipleLeadsHandler",
				"method":     r.Method,
				"url":        r.URL.Path,
				"remoteAddr": r.RemoteAddr,
				"leadID":     req.ID,
			}, fmt.Sprintf("Error updating lead with id %s: %v", req.ID, err))
			utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("error updating lead with id %s: %v", req.ID, err),
			})
			return
		}
		updatedLeads = append(updatedLeads, *updatedLead)
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "leads updated successfully",
		"leads":   updatedLeads,
	})
}

func CreateSingleLeadHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info(logrus.Fields{
		"action":     "CreateSingleLeadHandler",
		"method":     r.Method,
		"url":        r.URL.Path,
		"remoteAddr": r.RemoteAddr,
	}, "POST request received at /lead")

	var lead schemas.Lead

	err := utils.ParseBody(r, &lead)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "CreateSingleLeadHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Error decoding request body: %v", err))
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = utils.ValidateLead(&lead)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "CreateMultipleLeadsHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Validation failed: %v", err))
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	createdLead, err := models.CreateLead(&lead)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "CreateMultipleLeadsHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Error creating lead: %v", err))
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "lead created successfully",
		"lead":    createdLead,
	})
}

func GetRandomLeadHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info(logrus.Fields{
		"action":     "GetRandomLeadHandler",
		"method":     r.Method,
		"url":        r.URL.Path,
		"remoteAddr": r.RemoteAddr,
	}, "GET request received at /lead")

	lead, err := models.GetRandomLead()
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "GetRandomLeadHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Error retrieving random lead: %v", err))
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "random lead fetched successfully",
		"lead":    lead,
	})
}

func GetLeadByIDHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info(logrus.Fields{
		"action":     "GetLeadByIDHandler",
		"method":     r.Method,
		"url":        r.URL.Path,
		"remoteAddr": r.RemoteAddr,
	}, "GET request received at /lead/{id}")

	vars := mux.Vars(r)
	id := vars["id"]

	lead, err := models.GetLeadByID(id)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "GetLeadByIDHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Unable to fetch lead: %v", err))
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "lead fetched successfully",
		"lead":    lead,
	})
}

func DeleteLeadByIDHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info(logrus.Fields{
		"action":     "DeleteLeadByID",
		"method":     r.Method,
		"url":        r.URL.Path,
		"remoteAddr": r.RemoteAddr,
	}, "DELETE request received at /lead/{id}")

	vars := mux.Vars(r)
	id := vars["id"]

	err := models.DeleteLeadByID(id)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "GetLeadByIDHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Unable to delete lead: %v", err))
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "lead deleted successfully",
	})
}

func UpdateLeadByIDHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info(logrus.Fields{
		"action":     "UpdateLeadByIDHandler",
		"method":     r.Method,
		"url":        r.URL.Path,
		"remoteAddr": r.RemoteAddr,
	}, "PUT request received at /lead/{id}")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "UpdateLeadByIDHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Error parsing UUID: %v", err))
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid UUID format"})
		return
	}

	var updatedData map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "UpdateLeadByIDHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Error decoding request body: %v", err))
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	updatedLead, err := models.UpdateLeadByID(id, updatedData)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "UpdateLeadByIDHandler",
			"method":     r.Method,
			"url":        r.URL.Path,
			"remoteAddr": r.RemoteAddr,
		}, fmt.Sprintf("Unable to update lead: %v", err))
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "lead updated successfully",
		"lead":    updatedLead,
	})
}
