package handlers

import (
	"fmt"
	"go-simple-crm-tool/api/models"
	"go-simple-crm-tool/api/schemas"
	"go-simple-crm-tool/pkg/utils"
	"net/http"

	"github.com/sirupsen/logrus"
)

func CreateLeadsHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info(logrus.Fields{
		"action":     "CreateLeadHandler",
		"method":     r.Method,
		"url":        r.URL.Path,
		"remoteAddr": r.RemoteAddr,
	}, "POST request received at /leads")

	var leads []schemas.Lead

	err := utils.ParseBody(r, &leads)
	if err != nil {
		utils.Warn(logrus.Fields{
			"action":     "CreateLeadHandler",
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
				"action":     "CreateLeadHandler",
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
				"action":     "CreateLeadHandler",
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

	leads, err := models.GetAllLeadsHandler()
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
