package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

type Config struct {
	timeout time.Duration
	maxAge  int
}

func NewConfig(timeoutSeconds int, maxAge int) *Config {
	if timeoutSeconds < 0 {
		timeoutSeconds = 30
	}
	if maxAge < 0 {
		maxAge = 5
	}

	return &Config{
		timeout: time.Duration(timeoutSeconds) * time.Second,
		maxAge:  maxAge,
	}
}

func (c *Config) Cors(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           c.maxAge,
	}).Handler(next)
}

func (c *Config) SetTimeout(timeInSeconds int) *Config {
	if timeInSeconds >= 0 {
		c.timeout = time.Duration(timeInSeconds) * time.Second
	}
	return c
}

func (c *Config) GetTimeout() time.Duration {
	return c.timeout
}

func (c *Config) SetMaxAge(ageInSeconds int) *Config {
	if ageInSeconds >= 0 {
		c.maxAge = ageInSeconds
	}
	return c
}

func (c *Config) GetMaxAge() int {
	return c.maxAge
}
