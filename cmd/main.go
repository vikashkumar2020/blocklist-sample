package main

import (
	"blocklist/internal/common/register"
	"blocklist/internal/config"

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

	router := gin.Default()
	register.Routes(router, serverConfig)

	if err := router.Run(":" + serverConfig.Port); err != nil {
		utils.LogFatal(err)
	}
	utils.LogInfo("server started")
}
