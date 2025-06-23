package routes

import (
	"Web_template/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	router.Use(logger.GinRecovery(true), logger.GinLogger())
	
	router.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})
	
	return router
}
