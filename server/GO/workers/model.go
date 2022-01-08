package workers

import (
	"appbar/bars"
	"appbar/common"
	"appbar/role_permissions"
	"appbar/users"
	"fmt"
	"gorm.io/gorm/clause"
)

type Worker struct {
	User users.UserModel       `gorm:"many2many:workers"`
	Bar  bars.BarModel         `gorm:"many2many:workers"`
	Role role_permissions.Role `gorm:"many2many:workers"`
}

func FindAllWorkersFromBar(id string) error {
	db := common.GetDB()
	var model []Worker
	err := db.Where(fmt.Sprintf("id = %q", id)).Find(&model).Error
	return err
}

func AssignWorkerToBar(data interface{}) error {
	db := common.GetDB()
	err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(data).Error
	return err
}

func DeleteWorkerFromBar(data *Worker) error {
	db := common.GetDB()
	err := db.Model(data).Association("Worker").Delete(data)
	return err
}
