package bars

import (
	"github.com/gin-gonic/gin"
)

type BarSerializer struct {
	C *gin.Context
	BarModel
}

type BarWorkSerializer struct {
	C *gin.Context
	BarRolModel
}

func (s *BarSerializer) Response() BarResponse {
	response := BarResponse{
		Name: s.Name,
		Slug: s.Slug,
		City: s.City,
		Rol: "owner",
	}
	return response
}

func (s *BarWorkSerializer) Response() BarResponse {
	response := BarResponse{
		Name: s.Name,
		Slug: s.Slug,
		City: s.City,
		Rol: s.Rol,
	}
	return response
}

type BarResponse struct {
	Name	string	`json:"name"`
	Slug	string	`json:"slug"`
	City	string	`json:"city"`
	Rol		string	`json:"rol"`
}

type BarsSerializer struct {
	C       *gin.Context
	Bars 	[]BarModel
}

type BarsWorkSerializer struct {
	C       *gin.Context
	Bars 	[]BarRolModel
}

func (s *BarsSerializer) Response() []BarResponse {
	response := []BarResponse{}
	for _, bar := range s.Bars {
		serializer := BarSerializer{s.C, bar}
		response = append(response, serializer.Response())
	}
	return response
}

func (s *BarsWorkSerializer) Response() []BarResponse {
	response := []BarResponse{}
	for _, bar := range s.Bars {
		serializer := BarWorkSerializer{s.C, bar}
		response = append(response, serializer.Response())
	}
	return response
}