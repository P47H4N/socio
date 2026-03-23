package helpers

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

func ValidateUsername(username string) (bool, error) {
	if len(username) < 5 || len(username) > 20 {
		return false, errors.New("Username must be between 5 and 20 characters.")
	}

	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9.]+$`)
	if !usernameRegex.MatchString(username) {
		return false, errors.New("Only letters, numbers, and dots are allowed.")
	}
	return true, nil
}

func AccountStatusCalculator(status string, stime time.Time) error {
	deletionTime := stime
	duration := time.Since(deletionTime)
	if duration.Minutes() < 60 {
		return fmt.Errorf("User %s %d minutes ago", status, int(duration.Minutes()))
	}
	if duration.Hours() < 24 {
		return fmt.Errorf("User %s %d hours ago", status, int(duration.Hours()))
	}
	days := int(duration.Hours() / 24)
	return fmt.Errorf("User %s %d days ago", status, days)
}
