package users

import (
	"github.com/gin-gonic/gin"
	"appbar/common"
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
	Photo  string `json:"photo"`
	Token  string `json:"token"`
	Superadmin bool `json:"superadmin"`
}

func (myuser *UserSerializer) Response() UserResponse {
	myUserModel := myuser.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Name:   myUserModel.Name,
		Photo:  myUserModel.Photo,
		Token:  common.GenToken(myUserModel.ID),
	}
	return user
}

func (myuser *UserSerializer) SuperResponse(token string) UserResponse {
	myUserModel := myuser.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Name:   myUserModel.Name,
		Photo:  myUserModel.Photo,
		Token:  token,
		Superadmin: true,
	}
	return user
}

