package config

import (
	"log"
	"strconv"

	"github.com/datarohit/go-dynamodb-crud-api/utils/env"
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
		log.Printf("Invalid value for %s: %s, using default value: %s", envName, value, defaultValue)
		num, _ = strconv.Atoi(defaultValue)
	}
	return num
}
