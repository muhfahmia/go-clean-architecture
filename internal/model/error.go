package model

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/muhfahmia/pkg/enum"
	"github.com/muhfahmia/pkg/utils"
)

type AppError interface {
	GetError() error
	GetErrorMessage() string
	GetErrorType() enum.ErrorType
	GetErrorValidation() []appValidationError
}

type appError struct {
	Type            enum.ErrorType
	Message         string
	ErrorSource     error
	ErrorValidation []appValidationError
}

type appValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewAppError(err error, errType enum.ErrorType, message string) AppError {
	return &appError{
		ErrorSource: err,
		Type:        errType,
		Message:     message,
	}
}

func NewAppErrorValidation(err error, errType enum.ErrorType) AppError {
	appError := appError{
		ErrorSource: err,
		Type:        errType,
		Message:     "A few fields need attention—please review the highlights and try again!",
	}
	appError.SetErrorValidation()
	return &appError
}

func (e *appError) Error() string {
	return fmt.Sprintf("ErrSource: %s", e.ErrorSource.Error())
}

func (e *appError) GetError() error {
	return e.ErrorSource
}

func (e *appError) GetErrorMessage() string {
	return e.Message
}

func (e *appError) GetErrorType() enum.ErrorType {
	return e.Type
}

func (e *appError) GetErrorValidation() []appValidationError {
	return e.ErrorValidation
}

func (e *appError) SetErrorValidation() {
	verros := []appValidationError{}
	switch e.Type {
	case enum.ErrorValidation:
		errors, _ := e.ErrorSource.(validator.ValidationErrors)
		for _, err := range errors {
			var errorMessage string
			switch err.Tag() {
			// Basic validations
			case "required":
				errorMessage = fmt.Sprintf("%s is required", utils.HumanizeFieldName(err.Field()))
			case "required_if", "required_unless", "required_with", "required_without", "required_without_all":
				errorMessage = fmt.Sprintf("%s is required under these conditions", utils.HumanizeFieldName(err.Field()))

			// String validations
			case "min":
				if err.Type().Kind() == reflect.String {
					errorMessage = fmt.Sprintf("%s must be at least %s characters", utils.HumanizeFieldName(err.Field()), err.Param())
				} else {
					errorMessage = fmt.Sprintf("%s must be at least %s", utils.HumanizeFieldName(err.Field()), err.Param())
				}
			case "max":
				if err.Type().Kind() == reflect.String {
					errorMessage = fmt.Sprintf("%s cannot exceed %s characters", utils.HumanizeFieldName(err.Field()), err.Param())
				} else {
					errorMessage = fmt.Sprintf("%s cannot exceed %s", utils.HumanizeFieldName(err.Field()), err.Param())
				}
			case "len":
				errorMessage = fmt.Sprintf("%s must be exactly %s characters", utils.HumanizeFieldName(err.Field()), err.Param())

			// Number validations
			case "eq", "eq_ignore_case":
				errorMessage = fmt.Sprintf("%s must equal %s", utils.HumanizeFieldName(err.Field()), err.Param())
			case "ne", "ne_ignore_case":
				errorMessage = fmt.Sprintf("%s cannot equal %s", utils.HumanizeFieldName(err.Field()), err.Param())
			case "gt":
				errorMessage = fmt.Sprintf("%s must be greater than %s", utils.HumanizeFieldName(err.Field()), err.Param())
			case "gte":
				errorMessage = fmt.Sprintf("%s must be %s or greater", utils.HumanizeFieldName(err.Field()), err.Param())
			case "lt":
				errorMessage = fmt.Sprintf("%s must be less than %s", utils.HumanizeFieldName(err.Field()), err.Param())
			case "lte":
				errorMessage = fmt.Sprintf("%s must be %s or less", utils.HumanizeFieldName(err.Field()), err.Param())

			// Format validations
			case "email":
				errorMessage = "Please enter a valid email address"
			case "url":
				errorMessage = "Please enter a valid URL"
			case "uri":
				errorMessage = "Please enter a valid URI"
			case "uuid":
				errorMessage = "Please enter a valid UUID"
			case "uuid3", "uuid4", "uuid5":
				errorMessage = fmt.Sprintf("Please enter a valid UUID%s", strings.TrimPrefix(err.Tag(), "uuid"))
			case "alpha":
				errorMessage = fmt.Sprintf("%s can only contain letters", utils.HumanizeFieldName(err.Field()))
			case "alphanum":
				errorMessage = fmt.Sprintf("%s can only contain letters and numbers", utils.HumanizeFieldName(err.Field()))
			case "numeric":
				errorMessage = fmt.Sprintf("%s must be a numeric value", utils.HumanizeFieldName(err.Field()))
			case "hexadecimal":
				errorMessage = fmt.Sprintf("%s must be a hexadecimal number", utils.HumanizeFieldName(err.Field()))
			case "hexcolor":
				errorMessage = "Please enter a valid hex color code"
			case "rgb":
				errorMessage = "Please enter a valid RGB color"
			case "rgba":
				errorMessage = "Please enter a valid RGBA color"
			case "hsl":
				errorMessage = "Please enter a valid HSL color"
			case "hsla":
				errorMessage = "Please enter a valid HSLA color"
			case "e164":
				errorMessage = "Please enter a valid E.164 phone number"
			case "issn":
				errorMessage = "Please enter a valid ISSN"
			case "isbn":
				errorMessage = "Please enter a valid ISBN"
			case "isbn10":
				errorMessage = "Please enter a valid ISBN-10"
			case "isbn13":
				errorMessage = "Please enter a valid ISBN-13"

			// Date validations
			case "datetime":
				errorMessage = fmt.Sprintf("%s must be a valid date/time (format: %s)", utils.HumanizeFieldName(err.Field()), err.Param())
			case "timezone":
				errorMessage = "Please enter a valid timezone"

			// Custom validations
			case "msisdn":
				errorMessage = "Please enter a valid phone number starting with 08 or +62/62 (e.g., 0812345678 or +62812345678)"
			case "oneof":
				errorMessage = fmt.Sprintf("%s must be one of: %s", utils.HumanizeFieldName(err.Field()), strings.Replace(err.Param(), " ", ", ", -1))
			case "unique":
				errorMessage = fmt.Sprintf("%s contains duplicate values", utils.HumanizeFieldName(err.Field()))

			// File validations
			case "file":
				errorMessage = "Please provide a valid file"
			case "image":
				errorMessage = "Please provide a valid image file"
			case "username":
				errorMessage = "Please choose a valid username (letters, numbers, . or _ only, 1-30 characters)"

			// Default case
			default:
				errorMessage = fmt.Sprintf("%s contains an invalid value", utils.HumanizeFieldName(err.Field()))
			}
			verros = append(verros, appValidationError{
				Field:   err.Field(),
				Message: errorMessage,
			})
		}
		e.ErrorValidation = verros
	}

}
