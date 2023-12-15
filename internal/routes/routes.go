package routes

import (
	controller "blocklist/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterSpamRoutes(router *gin.RouterGroup) {
	router.POST("/spam", controller.CheckSpam)
}
