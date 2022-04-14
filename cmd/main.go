package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.186.1.2"})

	// Root
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
	})

	// Ping Test
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}
