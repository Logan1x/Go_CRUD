package main

import (
	"github.com/gin-gonic/gin"
	"github.com/logan1x/go_crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // default listen and serve on 0.0.0.0:8080
}
