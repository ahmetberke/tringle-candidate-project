package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello",
		})
	})

	PORT := os.Getenv("HTTP_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	err := router.Run(":" + PORT)
	if err != nil {
		panic(err)
	}
}
