package database

import (
	"blocklist/internal/config"
	"fmt"
)

func GenerateMongoConnectionString(config *config.DatabaseConfig) string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		config.User, config.Password, config.Host, config.Dbname)
}