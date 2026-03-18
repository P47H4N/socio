package helpers

import (
	"errors"
	"regexp"
)

func ValidateUsername(username string) (bool, error) {
	if len(username) < 3 || len(username) > 20 {
		return false, errors.New("Username must be between 3 and 20 characters.")
	}

	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9.]+$`)
	if !usernameRegex.MatchString(username) {
		return false, errors.New("Only letters, numbers, and dots are allowed.")
	}
	return true, nil
}
