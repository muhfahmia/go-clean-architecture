package config

import "github.com/go-playground/validator/v10"

func (c *appConfig) NewValidator() *validator.Validate {
	return validator.New()
}
