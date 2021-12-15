package bars

import (
	"appbar/common"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
	"github.com/gosimple/slug"
	// "github.com/gin-gonic/gin"
)

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
