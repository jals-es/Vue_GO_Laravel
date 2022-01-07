package products

import (
	"github.com/gin-gonic/gin"
	"appbar/common"
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

	// for _, thisType := range prodModelValidator.prodTypesModel.Types {
	// 	var model ProdTypeModel

	// 	model = thisType

	// 	if err := SaveOneType(model); err != nil {
	// 		my_error := common.NewError("prod", err)
	
	// 		c.JSON(http.StatusUnprocessableEntity, my_error)
	// 		return
	// 	}
	// }

	// for _, thisExtra := range prodModelValidator.prodExtrasModel.Extras {

	// 	var model ProdExtraModel

	// 	model = thisExtra

	// 	if err := SaveOneExtra(model); err != nil {
	// 		my_error := common.NewError("prod", err)
	
	// 		c.JSON(http.StatusUnprocessableEntity, my_error)
	// 		return
	// 	}
	// }

	c.JSON(http.StatusCreated, gin.H{"message": "Producto creado correctamente"})
}