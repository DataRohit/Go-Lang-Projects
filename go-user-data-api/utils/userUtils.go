package utils

import (
	"fmt"

	"github.com/datarohit/go-user-data-api/schemas"
)

func ValidateUser(user schemas.User) error {
	if user.Name == "" {
		return fmt.Errorf("name is required")
	}
	if user.Gender == "" {
		return fmt.Errorf("gender is required")
	}
	if user.Age <= 0 {
		return fmt.Errorf("age must be a positive integer")
	}
	return nil
}
