package env

import (
	"fmt"
	"os"
	"testing"
)

func TestGetEnvVariableOrExit(t *testing.T) {
	// Testing when the environment variable is set
	envVariable := "MY_ENV_VARIABLE"
	envVariableValue := "my_value"
	os.Setenv(envVariable, envVariableValue)

	result := GetEnvVariableOrExit(envVariable)
	if result != envVariableValue {
		t.Errorf("Expected %s, but got %s", envVariableValue, result)
	}

	// Testing when the environment variable is not set
	envVariable = "NOT_SET_ENV_VARIABLE"
	expectedErrorMessage := fmt.Sprintf("Env variable %s not set", envVariable)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic with message: %s", expectedErrorMessage)
		} else if r.(string) != expectedErrorMessage {
			t.Errorf("Expected panic with message: %s, but got: %s", expectedErrorMessage, r)
		}
	}()

	GetEnvVariableOrExit(envVariable)
}

func TestGetEnvVariableOrDefault(t *testing.T) {
	// Testing when the environment variable is empty and the default value is provided
	t.Run("Empty environment variable, default value provided", func(t *testing.T) {
		os.Setenv("ENV_VAR", "")
		defaultValue := "default"
		expected := defaultValue
		actual := GetEnvVariableOrDefault("ENV_VAR", defaultValue)
		if actual != expected {
			t.Errorf("Expected: %s, but got: %s", expected, actual)
		}
	})

	// Testing when the environment variable is set and the default value is not provided
	t.Run("Environment variable set, default value not provided", func(t *testing.T) {
		os.Setenv("ENV_VAR", "value")
		expected := "value"
		actual := GetEnvVariableOrDefault("ENV_VAR", "")
		if actual != expected {
			t.Errorf("Expected: %s, but got: %s", expected, actual)
		}
	})

	// Testing when the environment variable is set and the default value is provided
	t.Run("Environment variable set, default value provided", func(t *testing.T) {
		os.Setenv("ENV_VAR", "value")
		defaultValue := "default"
		expected := "value"
		actual := GetEnvVariableOrDefault("ENV_VAR", defaultValue)
		if actual != expected {
			t.Errorf("Expected: %s, but got: %s", expected, actual)
		}
	})

	// Testing when the environment variable is not set and the default value is not provided
	t.Run("Environment variable not set, default value not provided", func(t *testing.T) {
		os.Unsetenv("ENV_VAR")
		expected := ""
		actual := GetEnvVariableOrDefault("ENV_VAR", "")
		if actual != expected {
			t.Errorf("Expected: %s, but got: %s", expected, actual)
		}
	})
}
