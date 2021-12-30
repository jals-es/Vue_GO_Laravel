package role_permissions

import (
	"appbar/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRole(c *gin.Context) {
	roleValidator := NewRoleValidator()

	if err := roleValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := SaveOne(&roleValidator.roleModel); err != nil {
		roleError := common.NewError("role", err)
		c.JSON(http.StatusUnprocessableEntity, roleError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Role Created"})
}

func ListRoles(c *gin.Context) {
	data, _ := FindRoles()
	c.JSON(http.StatusAccepted, gin.H{"roles": data})
}

func AssignPermissionToRole(c *gin.Context) {

}
