package users

import (
	"github.com/gin-gonic/gin"
)

type ProfileSerializer struct {
	C *gin.Context
	UserModel
}

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Photo  string `json:"photo"`
	Passwd string `json:"passwd"`
	Status int    `json:"status"`
}

func (myuser *UserSerializer) Response() UserResponse {
	myUserModel := myuser.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Name:   myUserModel.Name,
		Email:  myUserModel.Email,
		Photo:  myUserModel.Photo,
		Passwd: myUserModel.Passwd,
		Status: myUserModel.Status,
	}
	return user
}
