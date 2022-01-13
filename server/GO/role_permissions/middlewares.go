package role_permissions

import (
	"appbar/common"
	"appbar/workers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PermissionMiddleware(c *gin.Context) {
	db := common.GetDB()

	userId := c.MustGet("my_user_id")
	barId := c.Request.Header["Bar"][0]
	permissionId := c.Request.Header["Permission"][0]

	var workerModel workers.Worker
	var integer int64

	db.Where("bar_id = ? and user_id = ?", barId, userId).Find(&workerModel)
	db.Preload("Permissions").Where("id = ?", permissionId).Count(&integer)

	if integer == 0 {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Not Allowed"})
		return
	}

	c.Next()
}
