package config

import (
	"blocklist/internal/utils"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port              string
	ServerApiPrefixV1 string
	BasePath          string
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Dbname   string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:              os.Getenv("PORT"),
		ServerApiPrefixV1: os.Getenv("SERVER_API_PREFIX_V1"),
		BasePath:          os.Getenv("SERVER_BASE_PATH"),
	}
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
	}
}

// LoadEnv loads environment variables from the .env
func LoadEnv() {

	loadEnvError := godotenv.Load(".env")
	if loadEnvError != nil {
		utils.LogFatal(loadEnvError)
	}
}
