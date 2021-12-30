package router

import (
	"appbar/bars"
	"appbar/role_permissions"
	"appbar/users"
	"github.com/gin-gonic/gin"
)

func GeneralRouter(router *gin.RouterGroup) {
	users.UsersRoutes(router.Group("/user"))
	bars.BarsRoutes(router.Group("/bar"))
	role_permissions.PrRoutes(router.Group("/pr"))
}
