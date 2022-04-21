package example

import (
	"github.com/gin-gonic/gin"
)

func ExampleController(router *gin.Engine) {
	exampleRouter := router.Group("/example")

	exampleRouter.GET("/", func(ctx *gin.Context) {
		status, data := GetExampleHelloWorld()
		ctx.JSON(status, data)
	})
}