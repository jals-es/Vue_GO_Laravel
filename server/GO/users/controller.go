package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	userModelValidator := NewUserModelValidator()

	// jsonData, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(jsonData)

	if err := userModelValidator.Bind(c); err != nil {
		// c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		fmt.Println("yee entra 1")
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
