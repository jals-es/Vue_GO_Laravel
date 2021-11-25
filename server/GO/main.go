package main

import (
	"appbar/router"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func main() {
	r := gin.Default()

	//Los nombres de las funciones en mayuscula, PERFAVOR
	router.GeneralRouter(r.Group("/api"))

	home := r.Group("/")

	fmt.Println(uuid.NewV4())

	home.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to APPBAR API REST",
		})
	})

	r.Run(":3000")
}
