package q

import (
	"fmt"
	"os"
	"strconv"
)

// ReadEnv retrieves the value of the environment variable named by the key.
// If the environment variable is not set, or is empty, the defaultValue is returned.
func ReadEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}

// ReadEnvStrict retrieves the value of the environment variable named by the key.
// It panics if the environment variable is not set or is empty.
func ReadEnvStrict(key string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	panic(fmt.Sprintf("Environment variable %s is required", key))
}

// ReadEnvInt retrieves the integer value of the environment variable named by the key.
// If the environment variable is not set, or is empty, or cannot be converted to an integer, the defaultValue is returned.
func ReadEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}

// ReadEnvIntStrict retrieves the integer value of the environment variable named by the key.
// It panics if the environment variable is not set, is empty, or cannot be converted to an integer.
func ReadEnvIntStrict(key string) int {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	panic(fmt.Sprintf("Environment variable %s is required", key))
}

// ReadEnvBool retrieves the boolean value of the environment variable named by the key.
// If the environment variable is not set, or is empty, or cannot be converted to a boolean, the defaultValue is returned.
func ReadEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}
	return defaultValue
}

// ReadEnvBoolStrict retrieves the boolean value of the environment variable named by the key.
// It panics if the environment variable is not set, is empty, or cannot be converted to a boolean.
func ReadEnvBoolStrict(key string) bool {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}
	panic(fmt.Sprintf("Environment variable %s is required", key))
}
