package cli

import (
	"bufio"
	"os"
	"strings"
)

// Retrieve environment variable with optional default value
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}

// Set environment variable when was not set yet
// You can optionally force setting the value even if value already exists
func SetEnv(key string, value string, force bool) error {
	existing := os.Getenv(key)
	if existing == "" || force {
		return os.Setenv(key, value)
	}

	return nil
}

// Unset environment variable if it exists
func UnsetEnv(key string) error {
	if _, exists := os.LookupEnv(key); exists {
		return os.Unsetenv(key)
	}

	return nil
}

// Read from environment variable or ask for it with optional default value
func ReadEnv(key string, question string, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" && question != "" {

		if defaultValue != "" {
			Printf(ColorNotice, "%s (%s) ", question, defaultValue)
		} else {
			Printf(ColorNotice, "%s ", question)
		}

		reader := bufio.NewReader(os.Stdin)
		value, _ = reader.ReadString('\n')
		value = strings.Replace(value, "\n", "", -1)

	}

	if value == "" {
		value = defaultValue
	}

	return value
}
