package main

import (
	"appbar/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Los nombres de las funciones en mayuscula, PERFAVOR
	router.GeneralRouter(r.Group("/api"))

	home := r.Group("/")

	home.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to APPBAR API REST",
		})
	})

	r.Run(":3000")
}
