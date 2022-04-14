package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupServer() *gin.Engine {
	server := gin.Default()

	server.SetTrustedProxies([]string{"192.186.1.2"})

	// Json Response
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
		fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
	})

	// String Respone, Ping Test
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	return server
}

func LaunchServer(PORT string) {
	server := setupServer()
	server.Run(PORT)
}
