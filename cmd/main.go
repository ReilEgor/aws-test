package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hi, I'm a simple API built with Go and Gin!",
			"status":  "success",
			"data":    nil,
		})
	})
	r.Run(":8080")
}
