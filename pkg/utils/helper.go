package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Cost factor menentukan kompleksitas hashing (4-31)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil // Password doesn't match, but not an error case
		}
		// Return other errors (malformed hash, wrong hash format, etc.)
		return false, err
	}

	return true, nil
}

func ReplaceMessageToReadable(field string) string {
	// toHumanReadable converts field names to more readable format
	// e.g., "FirstName" -> "First name", "user_id" -> "User ID"

	// Convert camelCase to space separated
	field = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(field, "${1} ${2}")
	// Convert snake_case to space separated
	field = strings.ReplaceAll(field, "_", " ")
	// Capitalize first letter
	if len(field) > 0 {
		field = strings.ToUpper(field[:1]) + field[1:]
	}
	return field
}

func HumanizeFieldName(field string) string {
	// Convert camelCase to "camel case"
	field = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(field, "${1} ${2}")
	field = strings.ToLower(field)

	// Special cases
	switch field {
	case "url":
		return "URL"
	case "uuid":
		return "UUID"
	case "id":
		return "ID"
	default:
		return field
	}
}

func GenerateSHA1Token() string {
	// Create a new random source with current time as seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a unique input combining timestamp and random number
	input := fmt.Sprintf("%d%d", time.Now().UnixNano(), r.Intn(1000000))

	// Generate SHA-1 hash
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)

	// Convert to hexadecimal string
	return hex.EncodeToString(hash)
}

func GenerateSHA256Token() string {
	// Create a new random source with current time as seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a unique input combining timestamp and random number
	input := fmt.Sprintf("%d%d", time.Now().UnixNano(), r.Intn(1000000))

	// Generate SHA-1 hash
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)

	// Convert to hexadecimal string
	return hex.EncodeToString(hash)
}

func TimeAsiaJakartaOnUTC() time.Time {
	return time.Now().Add(time.Hour * 7).UTC()
}
