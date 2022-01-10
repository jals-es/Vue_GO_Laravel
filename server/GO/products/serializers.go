package products

import (
	// "github.com/gin-gonic/gin"
)

type ProdExtraResponse struct {
	Name	string	`json:"name"`
	Descr	string	`json:"descr"`
	Price	float32	`json:"price"`
}

type ProdTypeResponse struct {
	Name	string	`json:"name"`
	Descr	string	`json:"descr"`
	Price	float32	`json:"price"`
}

type ProdResponse struct {
	Name	string	`json:"name"`
	Descr	string	`json:"descr"`
	Photo	string	`json:"photo"`
	Slug	string	`json:"slug"`
	Types	[]ProdTypeResponse `json:"types"`
	Extras	[]ProdExtraResponse	`json:"extras"`
}

type AllProdsResponse struct{
	Prods []AllProdsModel
}

type AllProdsSerializer struct{
	AllProdsResponse
}

func (s *AllProdsResponse) Response() []ProdResponse {
	response := []ProdResponse{}
	for _, prod := range s.Prods {
		serializer := ProdResponse{
			Name: prod.Name,
			Descr: prod.Descr,
			Photo: prod.Photo,
			Slug: prod.Slug,
		}

		for _, thisType := range prod.Types {
			serializer.Types = append(serializer.Types, ProdTypeResponse{
				Name: thisType.Name,
				Descr: thisType.Descr,
				Price: thisType.Price,
			})
		}

		for _, extra := range prod.Extras {
			serializer.Extras = append(serializer.Extras, ProdExtraResponse{
				Name: extra.Name,
				Descr: extra.Descr,
				Price: extra.Price,
			})
		}

		response = append(response, serializer)
	}

	return response
}