package role_permissions

import (
	"appbar/common"
	"appbar/workers"
	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(c *gin.Context) {
	db := common.GetDB()

	userId := c.MustGet("my_user_id")
	barId := c.Request.Header["Bar"][0]
	permissionName := c.Request.Header["Permission"][0]

	var workerModel workers.Worker
	var roleModel Role

	db.Where("bar_id = ? and user_id = ?", barId, userId).Find(&workerModel)
	db.Preload("Permissions").Where("id = ?", workerModel.Role).First(&roleModel)

}
