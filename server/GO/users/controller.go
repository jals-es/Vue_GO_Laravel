package users

import (
	"net/http"
	"appbar/common"
	"errors"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	userModelValidator := NewUserModelValidator()

	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.Set("my_user_model", userModelValidator.userModel)
	serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func Login(c *gin.Context){
	userModelValidator := NewLoginUserModelValidator()

	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	userModel, err := FindOne(&UserModel{Email: userModelValidator.userModel.Email})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	if userModel.checkPassword(userModelValidator.User.Passwd) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	// c.Set("my_user_model", userModel)
	UpdateContextUserModel(c, userModel.ID)
	serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}