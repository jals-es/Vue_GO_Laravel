package router

import (
	"appbar/users"

	"github.com/gin-gonic/gin"
)

func GeneralRouter(router *gin.RouterGroup) {
	users.UsersRoutes(router.Group("/user"))
}
