package bars

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func CreateBar(c *gin.Context){
	barModelValidator := NewBarModelValidator()

	if err := barModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Error en la validacion"})
		return
	}

	if err := SaveOne(&barModelValidator.barModel); err != nil {
		fmt.Println("entra2")
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Bar creado correctamente"})
}