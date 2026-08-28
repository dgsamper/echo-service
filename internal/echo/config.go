package echo

import (
	"errors"
	"strconv"
)

// Validates a PORT environment value
func ParsePort(value string) (string, error) {
	if value == "" {
		return "8080", nil
	}

	port, err := strconv.Atoi(value)
	if err != nil || port < 1 || port > 65535 {
		return "", errors.New("PORT must be a number between 1 and 65535")
	}
	return strconv.Itoa(port), nil
}
