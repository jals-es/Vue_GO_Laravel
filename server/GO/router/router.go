package router

import (
	"github.com/gin-gonic/gin"
	"appbar/products"
)

func GeneralRouter(router *gin.RouterGroup) {
	products.ProductRoutes(router.Group("/prod"))
}
