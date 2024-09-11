package config

import (
	"strconv"

	"github.com/datarohit/go-dynamodb-crud-api/utils/env"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"go.uber.org/zap"
)

type Config struct {
	Port        int
	Timeout     int
	Dialect     string
	DatabaseURI string
}

func GetConfig() Config {
	return Config{
		Port:        parseEnvToInt("PORT", "8080"),
		Timeout:     parseEnvToInt("TIMEOUT", "30"),
		Dialect:     env.GetEnv("DIALECT", "sqlite3"),
		DatabaseURI: env.GetEnv("DATABASE_URI", ":memory:"),
	}
}

func parseEnvToInt(envName, defaultValue string) int {
	value := env.GetEnv(envName, defaultValue)
	num, err := strconv.Atoi(value)
	if err != nil {
		logger.GetLogger().Warn("Invalid environment variable value",
			zap.String("variable", envName),
			zap.String("value", value),
			zap.String("default_value", defaultValue),
			zap.Error(err),
		)
		num, _ = strconv.Atoi(defaultValue)
	}
	return num
}
