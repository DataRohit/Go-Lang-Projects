package utils

import (
	"fmt"
	"go-simple-crm-tool/api/schemas"
)

func ValidateLead(lead *schemas.Lead, fieldsToValidate ...string) error {
	validateMap := make(map[string]bool)
	for _, field := range fieldsToValidate {
		validateMap[field] = true
	}

	if len(fieldsToValidate) == 0 {
		validateMap["FirstName"] = true
		validateMap["LastName"] = true
		validateMap["Email"] = true
		validateMap["Phone"] = true
		validateMap["Company"] = true
	}

	if validateMap["FirstName"] && lead.FirstName == "" {
		return fmt.Errorf("first_name is required")
	}
	if validateMap["LastName"] && lead.LastName == "" {
		return fmt.Errorf("last_name is required")
	}
	if validateMap["Email"] && lead.Email == "" {
		return fmt.Errorf("email is required")
	}
	if validateMap["Phone"] && lead.Phone == "" {
		return fmt.Errorf("phone is required")
	}
	if validateMap["Company"] && lead.Company == "" {
		return fmt.Errorf("company is required")
	}

	return nil
}
