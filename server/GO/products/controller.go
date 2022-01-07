package products

import (
	"github.com/gin-gonic/gin"
	// "fmt"
	"net/http"
)
 
func CreateProds(c *gin.Context){
	var prodModelValidator NewProdModelValidator

	if err := prodModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Error en la validacion"})
		return
	}	

	c.JSON(http.StatusCreated, gin.H{"message": "Producto creado correctamente"})
}