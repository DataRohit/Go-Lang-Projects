package models

import (
	"fmt"
	"go-simple-crm-tool/api/schemas"
	"go-simple-crm-tool/internal/database"
	"go-simple-crm-tool/pkg/utils"

	"github.com/sirupsen/logrus"
)

func CreateLead(lead *schemas.Lead) (*schemas.Lead, error) {
	result := database.DatabaseConnection.Create(lead)
	if result.Error != nil {
		utils.Warn(logrus.Fields{"action": "CreateLead"}, fmt.Sprintf("Error creating lead: %v", result.Error))
		return nil, result.Error
	}

	return lead, nil
}

func GetAllLeadsHandler() ([]schemas.Lead, error) {
	var leads []schemas.Lead

	result := database.DatabaseConnection.Find(&leads)
	if result.Error != nil {
		utils.Warn(logrus.Fields{"action": "GetAllLeadsHandler"}, fmt.Sprintf("Error retrieving stocks: %v", result.Error))
		return nil, result.Error
	}

	return leads, nil
}
