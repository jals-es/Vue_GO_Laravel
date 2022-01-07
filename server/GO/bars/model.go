package bars

import (
	"appbar/common"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
	"github.com/gosimple/slug"
	// "fmt"
)

type BarRolModel struct {
	ID		uuid.UUID `gorm:"column:id;type:uuid;primary_key;"`
	Name	string `gorm:"column:name"`
	Slug	string `gorm:"column:slug;unique"`
	Descr	string `gorm:"column:descr"`
	Lat		string `gorm:"column:lat"`
	Lon		string `gorm:"column:lon"`
	City	string `gorm:"column:city"`
	Address	string `gorm:"column:address"`
	Status	int    `gorm:"column:status"`
	Owner	uuid.UUID `gorm:"column:owner"`
	Rol 	string `gorm:"column:rname"`
}

type BarModel struct {
	ID		uuid.UUID `gorm:"column:id;type:uuid;primary_key;"`
	Name	string `gorm:"column:name"`
	Slug	string `gorm:"column:slug;unique"`
	Descr	string `gorm:"column:descr"`
	Lat		string `gorm:"column:lat"`
	Lon		string `gorm:"column:lon"`
	City	string `gorm:"column:city"`
	Address	string `gorm:"column:address"`
	Status	int    `gorm:"column:status"`
	Owner	uuid.UUID `gorm:"column:owner"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (BarModel) TableName() string {
	return "bars"
}

func (BarRolModel) TableName() string {
	return "bars"
}

func (b *BarModel) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV4()

	b.ID = id
	b.Slug = slug.Make(b.Name)
	b.Status = 0

	return
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func GetBars(id uuid.UUID) ([]BarRolModel, error) {

	db := common.GetDB()
	var models []BarRolModel
	var wmodels []BarRolModel

	tx := db.Begin()

	tx.Select("*, 'owner' as rname").Where(BarRolModel{Owner: id}).Find(&models)
	tx.Raw("SELECT b.*, r.name as rname FROM workers w, roles r, bars b WHERE w.id_user = ? AND w.id_rol = r.id AND w.id_bar = b.id", id).Scan(&wmodels)

	err := tx.Commit().Error

	for _, bar := range wmodels {
		models = append(models, bar)
	}

	return models, err
}

func GetBarBySlug(thisSlug string) (BarModel, error){

	var model BarModel

	db := common.GetDB()

	err := db.Where("slug = ?", thisSlug).Take(&model).Error

	return model, err
}