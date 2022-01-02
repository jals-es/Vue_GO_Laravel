package role_permissions

import (
	"appbar/users"
	"github.com/gin-gonic/gin"
)

func PrRoutes(router *gin.RouterGroup) {
	router.Use(users.AuthMiddleware(false))
	router.POST("/new-role", CreateRole)
	router.GET("/roles", ListRoles)
	router.DELETE("/role/*id")
	router.POST("/new-permission", CreatePermission)
	router.GET("/permissions", ListPermissions)
	router.POST("/assignation")
	router.DELETE("/permission/*id", DeletePermission)
	router.GET("/permission/*id", GetDataFromPermission)

}
