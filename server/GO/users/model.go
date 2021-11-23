package users

import (
	"appbar/common"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID     uint   `gorm:"column:id;primary_key"`
	Name   string `gorm:"column:name"`
	Email  string `gorm:"column:email;unique_index"`
	Photo  string `gorm:"column:photo"`
	Passwd string `gorm:"column:passwd"`
	Status int    `gorm:"column:status"`
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
