package role_permissions

import (
	"appbar/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Role Controller Logic

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

func GetDataFromRole(c *gin.Context) {
	id := strings.ReplaceAll(c.Param("id"), "/", "")
	data, err := FindOneRole(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't get Role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"role": data})

}

func DisableRole(c *gin.Context) {
	id := strings.ReplaceAll(c.Param("id"), "/", "")

	if err := DeleteRole(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role has been disabled"})
}

// Permission Controller Logic

func CreatePermission(c *gin.Context) {
	prValidator := NewPermissionValidator()
	if err := prValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if err := SaveOne(&prValidator.permissionModel); err != nil {
		permError := common.NewError("error", err)
		c.JSON(http.StatusUnprocessableEntity, permError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Permission Created"})
}

func ListPermissions(c *gin.Context) {
	data, _ := FindPermissions()
	c.JSON(http.StatusOK, gin.H{"permissions": data})
	return
}

func GetDataFromPermission(c *gin.Context) {
	id := strings.ReplaceAll(c.Param("id"), "/", "")
	data, err := FindOnePermission(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't get permission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func DeletePermission(c *gin.Context) {
	id := strings.ReplaceAll(c.Param("id"), "/", "")

	if err := DeletePerm(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete Role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Role with id: %s has been deleted successfully", id)})
}

// Assignation Controller Logic

func AssignPermissionToRole(c *gin.Context) {
	roleValidator := NewRoleValidator()

	if err := roleValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	if err := Assignation(&roleValidator.roleModel); err != nil {
		roleError := common.NewError("role", err)
		c.JSON(http.StatusUnprocessableEntity, roleError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Permission asignation OK"})
}

func RemovePermissionFromRole(c *gin.Context) {
	roleValidator := NewRoleValidator()

	if err := roleValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	if err := DeleteAssignation(&roleValidator.roleModel); err != nil {
		roleError := common.NewError("role", err)
		c.JSON(http.StatusUnprocessableEntity, roleError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Permission removal OK"})
}
