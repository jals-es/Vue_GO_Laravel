package products

import (
	"github.com/gin-gonic/gin"
	"appbar/common"
	"appbar/bars"
	"appbar/users"
	// "fmt"
	"net/http"
)
 
func CreateProds(c *gin.Context){
	var prodModelValidator NewProdModelValidator

	if err := prodModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Error en la validacion"})
		return
	}
	
	if err := SaveOne(&prodModelValidator.prodModel); err != nil {
		my_error := common.NewError("prod", err)

		c.JSON(http.StatusUnprocessableEntity, my_error)
		return
	}

	if err := SaveTypes(prodModelValidator.prodTypesModel.Types); err != nil {
		my_error := common.NewError("prod_types", err)

		c.JSON(http.StatusUnprocessableEntity, my_error)
		return
	}

	if err := SaveExtras(prodModelValidator.prodExtrasModel.Extras); err != nil {
		my_error := common.NewError("prod_extras", err)

		c.JSON(http.StatusUnprocessableEntity, my_error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Producto creado correctamente"})
}

func GetProds(c *gin.Context){
	slug := c.Param("slug")

	thisBar, err := bars.GetBarBySlug(slug)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "El bar no existe"})
		return
	}

	myUserModel := c.MustGet("my_user_model").(users.UserModel)

	if thisBar.Owner != myUserModel.ID {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "No tienes permisos en este bar"})
		return
	}

	prods, err := GetAllProds(thisBar.ID)

	if err != nil {
		my_error := common.NewError("prod", err)

		c.JSON(http.StatusUnprocessableEntity, my_error)
		return
	}

	serializer := AllProdsResponse{prods}

	c.JSON(http.StatusOK, gin.H{"prods": serializer.Response()})
}