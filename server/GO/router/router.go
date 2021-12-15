package router

import (
	"appbar/users"
	"appbar/bars"

	"github.com/gin-gonic/gin"
)

func GeneralRouter(router *gin.RouterGroup) {
	users.UsersRoutes(router.Group("/user"))
	bars.BarsRoutes(router.Group("/bar"))
}
