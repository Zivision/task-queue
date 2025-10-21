package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	// Start GIN router
	r := gin.Default()

	// Define GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// return simple json
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server
	r.Run()
}
