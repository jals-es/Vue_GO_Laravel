package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	userModelValidator := NewUserModelValidator()

	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		// c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.Set("my_user_model", userModelValidator.userModel)
	serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}
