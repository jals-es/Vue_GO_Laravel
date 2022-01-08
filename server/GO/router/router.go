package router

import (
	"appbar/users"
	"appbar/bars"
	"appbar/role_permissions"
	"appbar/users"
	"appbar/products"
	"github.com/gin-gonic/gin"
)

func GeneralRouter(router *gin.RouterGroup) {
	users.UsersRoutes(router.Group("/user"))
	bars.BarsRoutes(router.Group("/bar"))
	products.ProdsRoutes(router.Group("/prod"))
	role_permissions.PrRoutes(router.Group("/pr"))
}
