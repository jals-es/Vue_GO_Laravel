package products

import (
	"github.com/gin-gonic/gin"
	"appbar/users"
)

func ProdsRoutes(router *gin.RouterGroup) {
	router.Use(users.AuthMiddleware(true))
	router.POST("/", CreateProds)
	router.GET("/:slug", GetProds)
	router.DELETE("/:slug", DeleteProd)
}
