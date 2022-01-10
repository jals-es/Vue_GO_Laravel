package role_permissions

import (
	"appbar/common"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.NewV4()
		r.Status = 0
	}
	return
}

func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.NewV4()
	}
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

// Global Save Function

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

//Role Logic

func FindRoles() ([]Role, error) {
	db := common.GetDB()
	var model []Role
	err := db.Preload("Permissions").Find(&model).Error
	return model, err
}

func FindOneRole(id string) (Role, error) {
	db := common.GetDB()
	var model Role
	err := db.Where("id = ?", id).First(&model).Error
	return model, err
}

func DeleteRole(id string) error {
	db := common.GetDB()
	var model Role
	err := db.Model(&model).Where("id = ?", id).Update("status", 1).Error
	return err
}

//Permission Logic

func FindOnePermission(id string) (Permission, error) {
	db := common.GetDB()
	var model Permission
	err := db.Where("id = ?", id).First(&model).Error
	return model, err
}

func FindPermissions() ([]Permission, error) {
	db := common.GetDB()
	var model []Permission
	err := db.Debug().Find(&model).Error
	return model, err
}

func DeletePerm(id string) error {
	db := common.GetDB()
	var model Permission
	err := db.Where("id = ?", id).Delete(&model).Error
	return err
}

//Assignation Logic

func Assignation(data interface{}) error {
	db := common.GetDB()
	err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(data).Error
	return err
}

func DeleteAssignation(data *Role) error {
	db := common.GetDB()
	err := db.Model(data).Association("Permissions").Delete(data.Permissions)
	return err
}
