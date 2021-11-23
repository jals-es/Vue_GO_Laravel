package users

import (
	"appbar/common"

	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	User struct {
		Name   string `form:"name" json:"name" binding:"alphanum,min=4,max=30"`
		Email  string `form:"email" json:"email" binding:"email"`
		Photo  string `form:"photo" json:"photo" binding:"url"`
		Passwd string `form:"passwd" json:"passwd" binding:"min=8,max=255"`
		Status int    `form:"status" json:"status" binding:""`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

func (user *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, user)
	if err != nil {
		return err
	}
	user.userModel.Name = user.User.Name
	user.userModel.Email = user.User.Email
	user.userModel.Status = user.User.Status
	user.userModel.Photo = user.User.Photo

	if user.User.Passwd != common.NBRandomPassword {
		user.userModel.setPassword(user.User.Passwd)
	}

	return nil
}

// You can put the default value of a Validator here
func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	//userModelValidator.User.Email ="w@g.cn"
	return userModelValidator
}
