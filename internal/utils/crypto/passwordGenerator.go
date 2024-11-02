package crypto

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers      = "0123456789"
	symbols      = "!@#$%^&*()-_+=<>?"
)

func GeneratePassword(length int, includeSymbols bool) (string, error) {

	if length < 4 {
		return "", errors.New("password length must be at least 4 characters")
	}
	// Build character pool based on options
	allChars := lowerLetters + upperLetters + numbers
	if includeSymbols {
		allChars += symbols
	}

	password := make([]byte, length)
	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}
		password[i] = allChars[randomIndex.Int64()]
	}
	return string(password), nil
}
