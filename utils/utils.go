package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// LoadConfigs loads the configuration for the application
func LoadConfigs() *Config {
	return &Config{
		DBUsername: getEnv("DB_USERNAME", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnv("DB_PORT", ""),
		DBName:     getEnv("DB_NAME", ""),
	}
}

// InitDB initializes the database connection
func InitDB(username, password, host, port, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Config represents the application configuration
type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func getEnv(key, fallbackValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallbackValue
}
