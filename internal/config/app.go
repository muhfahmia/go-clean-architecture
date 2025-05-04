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
	Run()
	ProvideConfig()
	NewApp() *fiber.App
	NewPostgreSQLDatabase() *gorm.DB
	NewValidator() *validator.Validate
	GetEnvirontmentApp() enum.AppEnv
	GetApp() *fiber.App
	GetPostgreSQLDatabase() *gorm.DB
	GetValidator() *validator.Validate
}
type appConfig struct {
	env       enum.AppEnv
	port      string
	app       *fiber.App
	db        *gorm.DB
	validator *validator.Validate
}

func Bootstrap() AppConfig {
	appConfig := appConfig{}
	appConfig.CheckEnvirontment()
	appConfig.ProvideConfig()
	return &appConfig
}

func (c *appConfig) CheckEnvirontment() {
	envApp := os.Getenv("APP_ENV")
	if envApp == enum.AppProduction.String() {
		c.env = enum.AppProduction
	} else if envApp == enum.AppStaging.String() {
		c.env = enum.AppStaging
	} else {
		godotenv.Load()
		c.env = enum.AppDevelopment
	}
}

func (c *appConfig) ProvideConfig() {
	c.port = ":" + os.Getenv("APP_PORT")
	c.app = c.NewApp()
	c.db = c.NewPostgreSQLDatabase()
	c.validator = c.NewValidator()
}

func (c *appConfig) Run() {
	c.app.Listen(c.port)
}

func (c *appConfig) NewApp() *fiber.App {
	return fiber.New()
}

func (c *appConfig) GetApp() *fiber.App {
	return c.app
}

func (c *appConfig) GetEnvirontmentApp() enum.AppEnv {
	return c.env
}
