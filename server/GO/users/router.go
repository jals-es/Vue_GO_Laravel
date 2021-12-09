package users

import (
	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.RouterGroup) {
	router.POST("/", Register)
	router.PUT("/", Login)
}
