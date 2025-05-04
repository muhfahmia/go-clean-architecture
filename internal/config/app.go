package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/muhfahmia/pkg/enum"
	"gorm.io/gorm"
)

type AppConfig interface {
	ProvideConfig() *Config
	NewApp() *fiber.App
	NewPostgreSQLDatabase() *gorm.DB
	NewValidator() *validator.Validate
}

type appConfig struct {
	env enum.AppEnv
}

func NewAppConfig() AppConfig {
	appConfig := appConfig{}

	envApp := os.Getenv("APP_ENV")
	if envApp == enum.AppProduction.String() {
		appConfig.env = enum.AppProduction
	} else if envApp == enum.AppStaging.String() {
		appConfig.env = enum.AppStaging
	} else {
		godotenv.Load()
		appConfig.env = enum.AppDevelopment
	}
	return &appConfig
}

func (c *appConfig) NewApp() *fiber.App {
	return fiber.New()
}

func (c *appConfig) ProvideConfig() *Config {
	return &Config{
		Port:      os.Getenv("APP_PORT"),
		App:       c.NewApp(),
		DB:        c.NewPostgreSQLDatabase(),
		Validator: c.NewValidator(),
	}
}

type Config struct {
	Port      string
	App       *fiber.App
	DB        *gorm.DB
	Validator *validator.Validate
}

func (c *Config) Run() {
	c.App.Listen(":" + c.Port)
}
