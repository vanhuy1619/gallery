package config

import (
	"errors"
	"regexp"
)

func IsEmailValid(e string) (bool, error) {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(e) {
		return false, errors.New("Invalid email format")
	}
	return true, nil
}
