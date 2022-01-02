package role_permissions

import (
	"appbar/common"
	"fmt"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID    `gorm:"column:id;type:uuid;primary_key;"`
	Name        string       `gorm:"column:name"`
	Status      int          `gorm:"column:status"`
	Permissions []Permission `gorm:"many2many:rol_permissions;"`
}

type Permission struct {
	ID   uuid.UUID `gorm:"column:id; type:uuid;primary_key;"`
	Desc string    `gorm:"column:descr"`
	Name string    `gorm:"column:name"`
	Role []Role    `gorm:"many2many:rol_permissions;"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.NewV4()
	r.Status = 0
	return
}

func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewV4()
	return
}

type Tabler interface {
	TableName() string
}

func (Role) TableName() string {
	return "roles"
}

func (Permission) TableName() string {
	return "permissions"
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindRoles() ([]Role, error) {
	db := common.GetDB()
	var model []Role
	err := db.Preload("Permissions").Find(&model).Error
	return model, err
}

func FindOneRole(condition interface{}) (Role, error) {
	db := common.GetDB()
	var model Role
	err := db.Where(condition).First(&model).Error
	return model, err
}

func DeleteRole(id string) error {
	db := common.GetDB()
	var model Role
	err := db.Model(&model).Where(fmt.Sprintf("id = %q", id)).Update("status", 1).Error
	return err
}

func FindOnePermission(id string) (Permission, error) {
	db := common.GetDB()
	var model Permission
	err := db.Where(fmt.Sprintf("id = %q", id)).First(&model).Error
	return model, err
}

func FindPermissions() ([]Permission, error) {
	db := common.GetDB()
	var model []Permission
	err := db.Find(&model).Error
	return model, err
}

func DeletePerm(id string) error {
	db := common.GetDB()
	var model Permission
	err := db.Where(fmt.Sprintf("id = %q", id)).Delete(&model).Error
	return err
}
