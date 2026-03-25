package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	fmt.Println("server started at port 3000")
	if err := router.Run(":3000"); err != nil {
		panic(err)
	}
}
