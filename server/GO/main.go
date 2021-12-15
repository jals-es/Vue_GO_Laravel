package main

import (
	"appbar/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		
		c.Next()

	} else {
        
		// Everytime we receive an OPTIONS request, 
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real 
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {
	r := gin.Default()

	r.Use(CORS)


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
