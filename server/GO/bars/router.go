package bars

import (
	"github.com/gin-gonic/gin"
	"appbar/users"
)

func BarsRoutes(router *gin.RouterGroup) {
	router.Use(users.AuthMiddleware(true))
	router.POST("/", CreateBar)
}
