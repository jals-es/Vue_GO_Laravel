package role_permissions

import (
	"appbar/common"
	"github.com/gin-gonic/gin"
)

type RoleValidator struct {
	Role struct {
		Name string `form:"name" json:"name" binding:"required"`
	} `json:"role"`
	roleModel RoleModel `json:"-"`
}

type PermissionValidator struct {
	Permission struct {
		Name        string `form:"name" json:"name" binding:"required"`
		Description string `form:"descr" json:"descr" binding:"required"`
	} `json:"permission"`
	permissionModel PermissionModel `json:"-"`
}

func (role *RoleValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, role)

	if err != nil {
		return err
	}

	role.roleModel.Name = role.Role.Name

	return nil
}

func (permission *PermissionValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, permission)

	if err != nil {
		return err
	}

	permission.permissionModel.Name = permission.Permission.Name
	permission.permissionModel.Desc = permission.Permission.Description

	return nil
}

func NewRoleValidator() RoleValidator {
	return RoleValidator{}
}

func NewPermissionValidator() PermissionValidator {
	return PermissionValidator{}
}
