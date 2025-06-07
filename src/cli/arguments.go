package cli

import (
	"slices"
	"strconv"
	"strings"
)

// Retrieve argument with given keys or in given index position
func Arg(args []string, keys string, defaultValue string) string {

	options := strings.Split(keys, ",")
	positional := map[string]string{}
	counter := 0

	for _, arg := range args {
		// Check for --key=value -key=value or key=value like format first
		for _, key := range options {
			if strings.HasPrefix(arg, key+"=") {
				return strings.TrimPrefix(arg, key+"=")
			}
		}

		// If none of the options match, then put value into positional list
		if !strings.HasPrefix(arg, "-") {
			positional[strconv.Itoa(counter)] = arg
			counter++
		}
	}

	// Check positional index format, valid only for numeric keys
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

// Retrieve boolean flag with given keys
func Flag(args []string, keys string, defaultValue bool) bool {

	options := strings.Split(keys, ",")

	// Check for --key or -key like format
	for _, arg := range args {
		if slices.Contains(options, arg) {
			return true
		}
	}

	return defaultValue
}

// Retrieve multiple values for arguments with given key
func Multiple(args []string, key string, separator string) []string {

	result := []string{}
	for {
		// Find next argument key details
		item := Arg(args, key, "")

		// Break when no more items found
		// When found, remove key details from arguments list
		// Also append to the list of items
		if item == "" {
			break
		} else {
			index := slices.Index(args, key+"="+item)
			args = slices.Delete(args, index, index+1)

			if separator == "" {
				result = append(result, item)
			} else {
				result = append(result, strings.Split(item, separator)...)
			}
		}
	}

	return result
}
