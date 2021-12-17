package bars

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
	"appbar/common"
	"appbar/users"
)

func CreateBar(c *gin.Context){
	barModelValidator := NewBarModelValidator()

	if err := barModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Error en la validacion"})
		return
	}

	if err := SaveOne(&barModelValidator.barModel); err != nil {

		my_error := common.NewError("bar", err)

		c.JSON(http.StatusUnprocessableEntity, my_error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Bar creado correctamente"})
}

func GetYourBars(c *gin.Context){
	myUserModel := c.MustGet("my_user_model").(users.UserModel)

	myBars, err := GetBars(myUserModel.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "No tienes bares en tu propiedad"})
	}

	serializer := BarsWorkSerializer{c, myBars}
	c.JSON(http.StatusOK, gin.H{"bars": serializer.Response()})
}