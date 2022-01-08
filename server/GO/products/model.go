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