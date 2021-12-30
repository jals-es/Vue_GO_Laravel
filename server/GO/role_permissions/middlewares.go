package role_permissions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PermissionMiddleware(c *gin.Context) {

	c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Not allowed due to insufficient permissions"})
}
