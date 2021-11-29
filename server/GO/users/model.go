package users

import (
	"appbar/common"
	"errors"

	"gorm.io/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"github.com/drexedam/gravatar"
)

type UserModel struct {
	ID     uuid.UUID `gorm:"column:id;type:uuid;primary_key;"`
	Name   string `gorm:"column:name"`
	Email  string `gorm:"column:email;uniqueIndex"`
	Photo  string `gorm:"column:photo"`
	Passwd string `gorm:"column:passwd"`
	Status string `gorm:"column:status"`
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV4()

	u.ID = id
	u.Photo = gravatar.New(u.Email).AvatarURL();
	u.Status = "0";
	return
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (UserModel) TableName() string {
	return "users"
}

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Passwd = string(passwordHash)
	return nil
}

// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
// func (u *UserModel) checkPassword(password string) error {
// 	bytePassword := []byte(password)
// 	byteHashedPassword := []byte(u.Passwd)
// 	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
// }

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
