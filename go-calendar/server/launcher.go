package server

import (
	"fmt"
	"go-calendar-practice/go-calendar-practice/go-calendar/controller"
	"go-calendar-practice/go-calendar-practice/go-calendar/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
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

	// service/controller Test
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, VideoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, VideoController.Save(ctx))
	})

	return server
}

func Launch(PORT string) {
	server := setupServer()
	server.Run(PORT)
}
