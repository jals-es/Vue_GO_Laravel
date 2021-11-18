package products

import (
	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to PRODUCTS",
		})
	})
}