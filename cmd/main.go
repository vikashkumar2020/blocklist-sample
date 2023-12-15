package main

import (
	"blocklist/internal/common/register"
	"blocklist/internal/config"
	"blocklist/internal/infra/database/mongodb"

	"blocklist/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the config
	config.LoadEnv()
	utils.LogInfo("env loaded")

	// Initialize the server
	serverConfig := config.NewServerConfig()
	utils.LogInfo("server config loaded")

	// Initialize the database
	databaseConfig := config.NewDatabaseConfig()
	utils.LogInfo("database config loaded")

	// Initialise the connection to the database
	conn := mongodb.GetInstance(databaseConfig)
	utils.LogInfo("database connection initialised")

	defer conn.Disconnect(nil)

	router := gin.Default()
	register.Routes(router, serverConfig)

	if err := router.Run(":" + serverConfig.Port); err != nil {
		utils.LogFatal(err)
	}
	utils.LogInfo("server started")
}
