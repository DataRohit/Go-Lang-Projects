package models

import (
	"fmt"
	"go-simple-crm-tool/api/schemas"
	"go-simple-crm-tool/internal/database"
	"go-simple-crm-tool/pkg/utils"
	"log"

	"github.com/google/uuid"
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

func GetAllLeads() ([]schemas.Lead, error) {
	var leads []schemas.Lead

	result := database.DatabaseConnection.Find(&leads)
	if result.Error != nil {
		utils.Warn(logrus.Fields{"action": "GetAllLeads"}, fmt.Sprintf("Error retrieving leads: %v", result.Error))
		return nil, result.Error
	}

	return leads, nil
}

func DeleteLead(id uuid.UUID) error {
	var lead schemas.Lead

	result := database.DatabaseConnection.Where("id = ?", id).First(&lead)
	if result.Error != nil {
		utils.Warn(logrus.Fields{"action": "DeleteLead"}, fmt.Sprintf("Error finding lead with id %s: %v", id, result.Error))
		return result.Error
	}

	result = database.DatabaseConnection.Delete(&lead)
	if result.Error != nil {
		utils.Warn(logrus.Fields{"action": "DeleteLead"}, fmt.Sprintf("Error deleting lead with id %s: %v", id, result.Error))
		return result.Error
	}

	return nil
}

func UpdateLead(id uuid.UUID, updatedData map[string]interface{}) (*schemas.Lead, error) {
	var lead schemas.Lead

	result := database.DatabaseConnection.Where("id = ?", id).First(&lead)
	if result.Error != nil {
		log.Printf("Error finding lead with id %s: %v", id, result.Error)
		return nil, result.Error
	}

	result = database.DatabaseConnection.Model(&lead).Updates(updatedData)
	if result.Error != nil {
		log.Printf("Error updating lead with id %s: %v", id, result.Error)
		return nil, result.Error
	}

	return &lead, nil
}
