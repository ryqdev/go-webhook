package main

import (
	"fmt"
	"net/http"
	"os/exec"

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
		unreal()
		c.JSON(http.StatusOK, gin.H{
			"message": "[OK]",
		})

	})
	r.Run(":8181")
}

func unreal() {
	cmd := exec.Command("make", "build")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing 'make build':", err)
		return
	}
	fmt.Println("Build successful")
}
