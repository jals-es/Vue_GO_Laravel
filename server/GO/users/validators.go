package users

import (
	"appbar/common"

	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	User struct {
		Name   string `form:"name" json:"name" binding:"required"`
		Email  string `form:"email" json:"email" binding:"required,email"`
		Passwd string `form:"passwd" json:"passwd" binding:"required,min=8,max=255"`
	} `json:"users"`
	userModel UserModel `json:"-"`
}


func (user *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, user)

	if err != nil {
		return err
	}
	user.userModel.Name = user.User.Name
	user.userModel.Email = user.User.Email

	if user.User.Passwd != common.NBRandomPassword {
		user.userModel.setPassword(user.User.Passwd)
	}

	return nil
}

// You can put the default value of a Validator here
func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}

	return userModelValidator
}

type LoginUserModelValidator struct {
	User struct {
		Email  string `form:"email" json:"email" binding:"required,email"`
		Passwd string `form:"passwd" json:"passwd" binding:"required,min=8,max=255"`
	} `json:"users"`
	userModel UserModel `json:"-"`
}

func NewLoginUserModelValidator() LoginUserModelValidator {
	userModelValidator := LoginUserModelValidator{}

	return userModelValidator
}

func (user *LoginUserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, user)

	if err != nil {
		return err
	}

	user.userModel.Email = user.User.Email

	return nil
}
