package env

import (
	"os"
	"strings"
)

func GetEnv(env, defaultValue string) string {
	environment := strings.TrimSpace(os.Getenv(env))
	if environment == "" {
		return defaultValue
	}
	return environment
}
