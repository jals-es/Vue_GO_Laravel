package workers

import (
	"appbar/common"
	"gorm.io/gorm/clause"
)

type Worker struct {
	User string `gorm:"column:user_id;type:string;primary_key;"`
	Bar  string `gorm:"column:bar_id;type:string;primary_key"`
	Role string `gorm:"column:role_id;type:string;primary_key;"`
}

func FindAllWorkersFromBar(id string) error {
	db := common.GetDB()
	var model []Worker
	err := db.Preload(clause.Associations).Where("bar_id = ?", id).Find(&model).Error
	return err
}

func AssignWorkerToBar(data *Worker) error {
	db := common.GetDB()
	err := db.Create(&data).Error
	return err
}

func DeleteWorkerFromBar(data *Worker) error {
	db := common.GetDB()
	err := db.Delete(data).Error
	return err
}
