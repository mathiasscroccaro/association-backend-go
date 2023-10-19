package env

import (
	"fmt"
	"os"
)

func GetEnvVariableOrExit(envVariable string) string {
	value := os.Getenv(envVariable)
	if value == "" {
		panic(fmt.Sprintf("Env variable %s not set", envVariable))
	}
	return value
}

func GetEnvVariableOrDefault(envVariable string, defaultValue string) string {
	value := os.Getenv(envVariable)
	if value == "" {
		return defaultValue
	}
	return value
}
