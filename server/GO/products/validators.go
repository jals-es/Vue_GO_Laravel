package products

import (
	"appbar/common"
	"appbar/bars"
	"appbar/users"
	"github.com/gin-gonic/gin"
	// "net/http"
	"fmt"
	"errors"
)

type ProdExtraValidator struct {
	Name   	string `form:"name" json:"name" binding:"required,min=8,max=255"`
	Descr  	string `form:"descr" json:"descr" binding:"required"`
	Price   float32 `form:"price" json:"price" binding:"required"`
}

type ProdTypeValidator struct {
	Name   	string `form:"name" json:"name" binding:"required,min=8,max=255"`
	Descr  	string `form:"descr" json:"descr" binding:"required"`
	Price   float32 `form:"price" json:"price" binding:"required"`
}

type NewProdModelValidator struct {
	BarSlug		string	`form:"barslug" json:"barslug" binding:"required"`
	Prod struct {
		Name   	string `form:"name" json:"name" binding:"required,min=8,max=255"`
		Descr  	string `form:"descr" json:"descr" binding:"required"`
		Catego	string `form:"catego" json:"catego" binding:"required"`
		Photo   string `form:"photo" json:"photo" binding:"required"`
		Types	[]ProdTypeValidator `form:"types" json:"types" binding:"required"`
		Extras  []ProdExtraValidator `form:"extras" json:"extras"`
	} `json:"prod"`
	prodModel ProdModel `json:"-"`
	prodTypesModel ProdTypesModel `json:"-"`
	prodExtrasModel ProdExtrasModel `json:"-"`
}

func (prod *NewProdModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, prod)

	if err != nil {
		return err
	}
	
	prod.prodModel.Name = prod.Prod.Name
	prod.prodModel.Descr = prod.Prod.Descr
	prod.prodModel.ID_category = prod.Prod.Catego
	prod.prodModel.Photo = prod.Prod.Photo
	
	bar, err := bars.GetBarBySlug(prod.BarSlug)

	// fmt.Println(bar)

	if err != nil {
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "El bar no existe"})
		return err
	}

	myUserModel := c.MustGet("my_user_model").(users.UserModel)

	fmt.Println(myUserModel.ID)
	fmt.Println(bar.Owner)

	if bar.Owner != myUserModel.ID {
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "No tienes permisos para crear productos en este bar"})
		err = errors.New("permisos: no tienes permisos")
		return err
	}

	prod.prodModel.ID_bar = bar.ID

	for _, mytype := range prod.Prod.Types {
		var thisType ProdTypeModel

		thisType.ID_prod = prod.prodModel.ID
		thisType.Name = mytype.Name
		thisType.Descr = mytype.Descr
		thisType.Price = mytype.Price

		prod.prodTypesModel.Types = append(prod.prodTypesModel.Types, thisType)
	}

	for _, extra := range prod.Prod.Extras {
		var thisExtra ProdExtraModel

		thisExtra.ID_prod = prod.prodModel.ID
		thisExtra.Name = extra.Name
		thisExtra.Descr = extra.Descr
		thisExtra.Price = extra.Price

		prod.prodExtrasModel.Extras = append(prod.prodExtrasModel.Extras, thisExtra)
	}

	return nil
}