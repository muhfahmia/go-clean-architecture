package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/muhfahmia/pkg/enum"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database interface {
	Connect()
	GetDatabase() *gorm.DB
}

type database struct {
	selector enum.Database
	db       *gorm.DB
}

func NewDatabase(selector enum.Database) Database {
	db := database{}
	db.selector = selector
	db.Connect()
	return &db
}

func (d *database) Connect() {
	switch d.selector {
	case enum.PostgresDB:
		d.db = newPostgreSQLDatabase()
	}
}

func (d *database) GetDatabase() *gorm.DB {
	return d.db
}

func (c *appConfig) NewPostgreSQLDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// Membuka koneksi
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Use the database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Set the maximum number of open connections
	sqlDB.SetMaxOpenConns(100)

	// Set the maximum number of idle connections
	sqlDB.SetMaxIdleConns(10)

	// Set the maximum lifetime of a connection
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping the database to verify the connection
	// err = sqlDB.Ping()
	// if err != nil {
	// 	log.Fatalf("Failed to ping database: %v", err)
	// }

	fmt.Println("Successfully connected to database.")

	return db
}

func (c *appConfig) GetPostgreSQLDatabase() *gorm.DB {
	return c.db
}

func newPostgreSQLDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// Membuka koneksi
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Use the database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Set the maximum number of open connections
	sqlDB.SetMaxOpenConns(100)

	// Set the maximum number of idle connections
	sqlDB.SetMaxIdleConns(10)

	// Set the maximum lifetime of a connection
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping the database to verify the connection
	// err = sqlDB.Ping()
	// if err != nil {
	// 	log.Fatalf("Failed to ping database: %v", err)
	// }

	fmt.Println("Successfully connected to database.")

	return db
}
