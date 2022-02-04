package users

import (
	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.RouterGroup) {
	router.Use(AuthMiddleware(false))
	router.POST("/", Register)
	router.PUT("/", Login)

	router.Use(AuthMiddleware(true))
	router.GET("/", CheckToken)
}
