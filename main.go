package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/webhook", func(c *gin.Context) {
		fmt.Printf("%s\n", "At least I got here")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})
	r.Run(":8181") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
