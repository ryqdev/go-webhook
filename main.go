package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello webhook",
		})
	})

	// ping is for testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// receive webhook
	r.POST("/webhook", func(c *gin.Context) {
		webhook()
		c.JSON(http.StatusOK, gin.H{
			"message": "[OK]",
		})

	})
	r.Run("0.0.0.0:1818")
}

func webhook() {
	cmd := exec.Command("make", "run")

	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing 'make build':", err)
		return
	}

	fmt.Println("Build successful")
}
