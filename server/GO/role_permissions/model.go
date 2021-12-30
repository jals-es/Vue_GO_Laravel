package role_permissions

import (
	"appbar/common"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type RoleModel struct {
	ID     uuid.UUID `gorm:"column:id;type:uuid;primary_key;"`
	Name   string    `gorm:"column:name"`
	Status int       `gorm:"column:status"`
}

type PermissionModel struct {
	ID   uuid.UUID `gorm:"column:id; type:uuid;primary_key;"`
	Desc string    `gorm:"column:descr"`
	Name string    `gorm:"column:name"`
}

type RolePermissionModel struct {
	IdRole       uuid.UUID `gorm:"column:id_rol; type:uuid; primary_key"`
	IdPermission uuid.UUID `gorm:"column:id_perm; type:uuid; primary_key"`
}

func (r *RoleModel) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.NewV4()
	r.Status = 0
	return
}

type Tabler interface {
	TableName() string
}

func (RoleModel) TableName() string {
	return "roles"
}

func (PermissionModel) TableName() string {
	return "permissions"
}

func (RolePermissionModel) TableName() string {
	return "rol_permissions"
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindRoles() ([]RoleModel, error) {
	db := common.GetDB()
	var model []RoleModel
	err := db.Find(&model).Error
	return model, err
}

func FindOneRole(condition interface{}) (RoleModel, error) {
	db := common.GetDB()
	var model RoleModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func FindOnePermission(condition interface{}) (PermissionModel, error) {
	db := common.GetDB()
	var model PermissionModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func FindPermissions() ([]PermissionModel, error) {
	db := common.GetDB()
	var model []PermissionModel
	err := db.Find(&model).Error
	return model, err
}

func GetPermissionFromRole(condition interface{}) (RolePermissionModel, error) {
	db := common.GetDB()
	var model RolePermissionModel
	err := db.Where(condition).First(&model).Error
	return model, err
}
