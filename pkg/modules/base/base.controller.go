package base

import (
	"github.com/gin-gonic/gin"
)

func AddController(router *gin.Engine) {
	exampleRouter := router.Group("/")

	exampleRouter.GET("/", func(ctx *gin.Context) {
		status, data := GetRootService()
		ctx.JSON(status, data)
	})
}
