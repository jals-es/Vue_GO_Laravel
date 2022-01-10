package workers

import (
	"appbar/users"
	"github.com/gin-gonic/gin"
)

func WorkerRoutes(router *gin.RouterGroup) {
	router.Use(users.AuthMiddleware(false))

	router.POST("/assignation", WorkerAssignation)
	router.POST("/removal", WorkerRemoval)
	router.GET("/bar/*id", GetWorkersFromBar)
}
