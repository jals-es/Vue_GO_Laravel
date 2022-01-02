package role_permissions

import (
	"appbar/users"
	"github.com/gin-gonic/gin"
)

func PrRoutes(router *gin.RouterGroup) {
	router.Use(users.AuthMiddleware(false))

	router.POST("/new-role", CreateRole)
	router.GET("/roles", ListRoles)
	router.DELETE("/role/*id", DisableRole)
	router.GET("/role/*id")

	router.POST("/new-permission", CreatePermission)
	router.GET("/permissions", ListPermissions)
	router.DELETE("/permission/*id", DeletePermission)
	router.GET("/permission/*id", GetDataFromPermission)

	router.POST("/assignation")
	router.DELETE("/assignation")
}
