package cli

import (
	"bufio"
	"os"
	"strings"
)

// Read from environment variable or ask for it
func Read(env string, question string, defaultValue string) string {

	value := os.Getenv(env)

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
