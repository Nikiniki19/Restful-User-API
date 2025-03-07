package utils

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(password string) bool {
	upperCase := regexp.MustCompile(`[A-Z]`)
	lowerCase := regexp.MustCompile(`[a-z]`)
	digit := regexp.MustCompile(`\d`)
	specialChar := regexp.MustCompile(`[@$!%*?&]`)

	return upperCase.MatchString(password) &&
		lowerCase.MatchString(password) &&
		digit.MatchString(password) &&
		specialChar.MatchString(password)
}

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error in hashing the password : %w", err)
	}
	return string(hashedPass), nil

}
