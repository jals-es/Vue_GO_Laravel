package users

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

type ProfileSerializer struct {
	C *gin.Context
	UserModel
}

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Photo  string `json:"photo"`
	Passwd string `json:"passwd"`
	Status string `json:"status"`
}

func (myuser *UserSerializer) Response() UserResponse {
	myUserModel := myuser.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		ID:     myUserModel.ID,
		Name:   myUserModel.Name,
		Email:  myUserModel.Email,
		Photo:  myUserModel.Photo,
		Passwd: myUserModel.Passwd,
		Status: myUserModel.Status,
	}
	return user
}
