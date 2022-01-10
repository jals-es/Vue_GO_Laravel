package bars

import (
	"appbar/common"
	"appbar/users"
	"github.com/gin-gonic/gin"
)

type BarModelValidator struct {
	Bar struct {
		Name    string `form:"name" json:"name" binding:"required"`
		Descr   string `form:"descr" json:"descr" binding:"required"`
		Lat     string `form:"lat" json:"lat" binding:"required"`
		Lon     string `form:"lon" json:"lon" binding:"required"`
		City    string `form:"city" json:"city" binding:"required"`
		Address string `form:"address" json:"address" binding:"required"`
	} `json:"bar"`
	barModel BarModel `json:"-"`
}

func (bar *BarModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, bar)

	if err != nil {
		return err
	}

	bar.barModel.Name = bar.Bar.Name
	bar.barModel.Descr = bar.Bar.Descr
	bar.barModel.Lat = bar.Bar.Lat
	bar.barModel.Lon = bar.Bar.Lon
	bar.barModel.City = bar.Bar.City
	bar.barModel.Address = bar.Bar.Address

	myUserModel := c.MustGet("my_user_model").(users.UserModel)

	bar.barModel.Owner = myUserModel.ID

	return nil
}

func NewBarModelValidator() BarModelValidator {
	barModelValidator := BarModelValidator{}

	return barModelValidator
}
