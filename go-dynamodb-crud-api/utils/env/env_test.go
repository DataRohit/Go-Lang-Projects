package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	t.Run("Should return default value when environment variable is not set", func(t *testing.T) {
		defaultValue := "GOLANG"
		environment := "PROGRAM"
		assert.Equal(t, defaultValue, GetEnv(environment, defaultValue))
	})

	t.Run("Should return the environment variable value if set", func(t *testing.T) {
		defaultValue := ""
		environment := "HOME"
		os.Setenv(environment, "/test/home")
		assert.Equal(t, "/test/home", GetEnv(environment, defaultValue))
		os.Unsetenv(environment)
	})

	t.Run("Should return default value when environment variable is empty", func(t *testing.T) {
		environment := "EMPTY_ENV"
		os.Setenv(environment, "")
		defaultValue := "DEFAULT"
		assert.Equal(t, defaultValue, GetEnv(environment, defaultValue))
		os.Unsetenv(environment)
	})
}
