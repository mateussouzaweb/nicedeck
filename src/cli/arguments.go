package cli

import (
	"strconv"
	"strings"
)

// Retrieve argument with given keys or in given index position
func Arg(args []string, keys string, defaultValue string) string {

	options := strings.Split(keys, ",")
	positional := map[string]string{}
	counter := 0

	// Check for --key=value or key=value like format first
	for _, arg := range args {
		for _, key := range options {
			if strings.HasPrefix(arg, key+"=") {
				return strings.TrimPrefix(arg, key+"=")
			}
		}
		if !strings.HasPrefix(arg, "--") {
			positional[strconv.Itoa(counter)] = arg
			counter++
		}
	}

	// Check positional index format, valid only for numbers
	for _, index := range options {
		if _, err := strconv.Atoi(index); err != nil {
			continue
		}
		if value, ok := positional[index]; ok {
			return value
		}
	}

	return defaultValue
}
