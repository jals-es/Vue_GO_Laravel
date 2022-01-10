package role_permissions

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type PermSerializer struct {
	C *gin.Context
	Permission
}

type RoleSerializer struct {
	C *gin.Context
	Role
}

type MultiPermSerializer struct {
	C           *gin.Context
	Permissions []Permission
}

type MultiRoleSerializer struct {
	C     *gin.Context
	Roles []Role
}

type RoleResponse struct {
	ID          uuid.UUID            `json:"id"`
	Name        string               `json:"name"`
	Status      int                  `json:"status"`
	Permissions []PermissionResponse `json:"permissions"`
}

type PermissionResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Desc string    `json:"desc"`
}

func (serial *PermSerializer) Response() PermissionResponse {
	response := PermissionResponse{
		ID:   serial.ID,
		Name: serial.Name,
		Desc: serial.Desc,
	}
	return response
}

func (serial *MultiPermSerializer) Response() []PermissionResponse {
	var response []PermissionResponse
	for _, permission := range serial.Permissions {
		serializer := PermSerializer{serial.C, permission}
		response = append(response, serializer.Response())
	}
	return response
}

func (serial *RoleSerializer) Response() RoleResponse {
	response := RoleResponse{
		ID:     serial.ID,
		Name:   serial.Name,
		Status: serial.Status,
	}
	response.Permissions = make([]PermissionResponse, 0)
	for _, permission := range serial.Permissions {
		serializer := PermSerializer{serial.C, permission}
		response.Permissions = append(response.Permissions, serializer.Response())
	}
	return response
}

func (serial *MultiRoleSerializer) Response() []RoleResponse {
	var response []RoleResponse
	for _, data := range serial.Roles {
		role := RoleResponse{
			ID:     data.ID,
			Name:   data.Name,
			Status: data.Status,
		}
		role.Permissions = make([]PermissionResponse, 0)
		for _, permission := range data.Permissions {
			perm := PermSerializer{serial.C, permission}
			role.Permissions = append(role.Permissions, perm.Response())
		}
		serializer := RoleSerializer{serial.C, data}
		response = append(response, serializer.Response())
	}
	return response
}
