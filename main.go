package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

//	type WebhookBody struct {
//		Command string `json:"command" binding:"required"`
//	}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/webhook", func(c *gin.Context) {
		//var cmd WebhookBody
		//c.BindJSON(&cmd)
		//fmt.Println(cmd.Command)
		webhook()
		c.JSON(http.StatusOK, gin.H{
			"message": "[OK]",
		})

	})
	r.Run("0.0.0.0:8181")
}

func webhook() {
	cmd := exec.Command("make", "mdbook")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing 'make build':", err)
		return
	}
	fmt.Println("Build successful")
}
