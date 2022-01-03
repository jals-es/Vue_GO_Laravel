package role_permissions

import (
	"appbar/common"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type RoleValidator struct {
	Role struct {
		ID          uuid.UUID    `form:"id" json:"id"`
		Name        string       `form:"name" json:"name" binding:"required"`
		Permissions []Permission `form:"permissions" json:"permissions"`
	} `json:"role"`
	roleModel Role `json:"-"`
}

type PermissionValidator struct {
	Permission struct {
		Name        string `form:"name" json:"name" binding:"required"`
		Description string `form:"descr" json:"descr" binding:"required"`
	} `json:"permission"`
	permissionModel Permission `json:"-"`
}

func (role *RoleValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, role)

	if err != nil {
		return err
	}
	role.roleModel.ID = role.Role.ID
	role.roleModel.Name = role.Role.Name
	role.roleModel.Permissions = role.Role.Permissions

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
