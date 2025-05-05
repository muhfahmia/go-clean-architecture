package config

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (c *appConfig) NewValidator() *validator.Validate {
	v := validator.New()
	// Register custom validation for msisdn_id
	err := v.RegisterValidation("msisdn", validateMSISDN)
	if err != nil {
		panic(err) // or handle error appropriately
	}

	// Add this for username validation
	err = v.RegisterValidation("username", validateUsername)
	if err != nil {
		panic(err)
	}

	// Add this for username validation
	err = v.RegisterValidation("identifier", validateIdentifier)
	if err != nil {
		panic(err)
	}

	// Add this for username validation
	err = v.RegisterValidation("latitude", validateLatitude)
	if err != nil {
		panic(err)
	}

	// Add this for username validation
	err = v.RegisterValidation("longitude", validateLongitude)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *appConfig) GetValidator() *validator.Validate {
	return c.validator
}

func validateMSISDN(fl validator.FieldLevel) bool {
	msisdn := fl.Field().String()

	// Remove all non-digit characters
	cleanMSISDN := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, msisdn)

	// Check length requirements
	if len(cleanMSISDN) < 8 || len(cleanMSISDN) > 15 {
		return false
	}

	// Convert international prefixes to local format
	if strings.HasPrefix(cleanMSISDN, "62") {
		cleanMSISDN = "0" + cleanMSISDN[2:]
	} else if strings.HasPrefix(cleanMSISDN, "621") {
		cleanMSISDN = "0" + cleanMSISDN[3:]
	}

	// Validate starts with 08
	if !strings.HasPrefix(cleanMSISDN, "08") {
		return false
	}

	return true
}

// Add this to your CustomValidator struct methods
func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	// Instagram username rules:
	// 1. 1-30 characters long
	// 2. Only letters, numbers, periods, underscores
	// 3. Can't start or end with period or underscore
	// 4. Can't have consecutive periods or underscores
	// 5. Can't be all numbers

	if len(username) < 1 || len(username) > 30 {
		return false
	}

	// Check allowed characters (letters, numbers, ., _)
	validChars := regexp.MustCompile(`^[a-zA-Z0-9._]+$`)
	if !validChars.MatchString(username) {
		return false
	}

	// Can't start/end with . or _
	if strings.HasPrefix(username, ".") || strings.HasPrefix(username, "_") ||
		strings.HasSuffix(username, ".") || strings.HasSuffix(username, "_") {
		return false
	}

	// No consecutive . or _
	if strings.Contains(username, "..") || strings.Contains(username, "__") ||
		strings.Contains(username, "._") || strings.Contains(username, "_.") {
		return false
	}

	// Can't be all numbers
	if _, err := strconv.Atoi(username); err == nil {
		return false
	}

	return true
}

func validateIdentifier(fl validator.FieldLevel) bool {
	identifier := fl.Field().String()
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	msisdnRegex := regexp.MustCompile(`^\+?[0-9]{10,15}$`)

	return usernameRegex.MatchString(identifier) || emailRegex.MatchString(identifier) || msisdnRegex.MatchString(identifier)
}

// ValidateLatitude checks if a field is a valid latitude (-90 to 90)
func validateLatitude(fl validator.FieldLevel) bool {
	lat, ok := fl.Field().Interface().(float64)
	if !ok {
		return false
	}
	return lat >= -90 && lat <= 90
}

// ValidateLongitude checks if a field is a valid longitude (-180 to 180)
func validateLongitude(fl validator.FieldLevel) bool {
	lng, ok := fl.Field().Interface().(float64)
	if !ok {
		return false
	}
	return lng >= -180 && lng <= 180
}
