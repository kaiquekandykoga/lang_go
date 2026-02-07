package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "okay",
		})
	})

	r.GET("/time", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": time.Now().Format(time.RFC3339),
		})
	})

	r.Run()
}
