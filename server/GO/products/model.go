package products

import (
	"appbar/common"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
	"github.com/gosimple/slug"
	// "fmt"
)

type ProdExtraModel struct {
	ID		uuid.UUID `gorm:"column:id;type:uuid;primary_key;"`
	ID_prod	uuid.UUID `gorm:"column:id_prod"`
	Name	string `gorm:"column:name"`
	Descr	string `gorm:"column:descr"`
	Price	float32`gorm:"column:price"`
	Status	int `gorm:"column:status"`
}

type ProdExtrasModel struct {
	Extras	[]ProdExtraModel
}

type ProdTypeModel struct {
	ID		uuid.UUID `gorm:"column:id;type:uuid;primary_key;"`
	ID_prod	uuid.UUID `gorm:"column:id_prod"`
	Name	string `gorm:"column:name"`
	Descr	string `gorm:"column:descr"`
	Price	float32`gorm:"column:price"`
	Status	int `gorm:"column:status"`
}

type ProdTypesModel struct {
	Types	[]ProdTypeModel
}

type ProdModel struct {
	ID			uuid.UUID `gorm:"column:id;type:uuid;primary_key;"`
	ID_bar		uuid.UUID `gorm:"column:id_bar"`
	ID_category	string `gorm:"column:id_category"`
	Name		string `gorm:"column:name"`
	Descr		string `gorm:"column:descr"`
	Status		int `gorm:"column:status"`
	Photo		string `gorm:"column:photo"`
	Slug		string `gorm:"column:slug"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (ProdModel) TableName() string {
	return "prods"
}

func (ProdTypeModel) TableName() string {
	return "prods_extras"
}

func (ProdExtraModel) TableName() string {
	return "prods_types"
}

func (p *ProdModel) BeforeCreate(tx *gorm.DB) (err error) {
	// id := uuid.NewV4()

	// p.ID = id
	p.Slug = slug.Make(p.Name)
	p.Status = 0

	return
}

func (t *ProdTypeModel) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV4()

	t.ID = id
	t.Status = 0

	return
}

func (e *ProdExtraModel) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV4()

	e.ID = id
	e.Status = 0

	return
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func SaveTypes(data []ProdTypeModel) error {
	db := common.GetDB()

	tx := db.Begin()
	
	for _, thisType := range data {
		tx.Save(&thisType)
	}

	err := tx.Commit().Error
	
	return err
}

func SaveExtras(data []ProdExtraModel) error {
	db := common.GetDB()

	tx := db.Begin()
	
	for _, thisExtra := range data {
		tx.Save(&thisExtra)
	}

	err := tx.Commit().Error
	
	return err
}

type AllProdsModel struct{
	ID			uuid.UUID
	ID_bar		uuid.UUID
	ID_category	string
	Name		string
	Descr		string
	Status		int
	Photo		string
	Slug		string
	Types		[]ProdTypeModel
	Extras		[]ProdExtraModel
}

func GetAllProds(id_bar uuid.UUID) ([]AllProdsModel, error){
	
	var response []AllProdsModel 
	var prods []ProdModel

	db := common.GetDB()
	tx := db.Begin()

	tx.Where("id_bar = ?", id_bar).Find(&prods)
	
	for _, prod := range prods {
		var newProd AllProdsModel

		newProd.ID = prod.ID
		newProd.ID_bar = prod.ID_bar
		newProd.ID_category = prod.ID_category
		newProd.Name = prod.Name
		newProd.Descr = prod.Descr
		newProd.Status = prod.Status
		newProd.Photo = prod.Photo
		newProd.Slug = prod.Slug

		var types []ProdTypeModel
		tx.Where("id_prod = ?", prod.ID).Find(&types)

		var extras []ProdExtraModel
		tx.Where("id_prod = ?", prod.ID).Find(&extras)

		newProd.Types = types
		newProd.Extras = extras

		response = append(response, newProd)
	}

	err := tx.Commit().Error

	return response, err
}

func GetProd(slug string) (ProdModel, error){

	var prod ProdModel

	db := common.GetDB()

	err := db.Where("slug = ?", slug).Find(&prod).Error

	return prod, err
}

func DelProd(prod ProdModel) error{ 

	db := common.GetDB()

	tx := db.Begin()

	tx.Delete(&prod)

	var types []ProdTypeModel
	tx.Where("id_prod = ?", prod.ID).Find(&types)
	tx.Delete(&types)

	var extras []ProdExtraModel
	tx.Where("id_prod = ?", prod.ID).Find(&extras)
	tx.Delete(&extras)

	err := tx.Commit().Error

	return err
}