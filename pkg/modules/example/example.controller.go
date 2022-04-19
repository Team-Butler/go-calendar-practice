package example

import (
	"github.com/gin-gonic/gin"
)

func ExampleController(r *gin.Engine) {
	exampleRouter := r.Group("/example")

	exampleRouter.GET("/", func(c *gin.Context) {
		status, data := GetExampleHelloWorld()
		c.JSON(status, data)
	})
}