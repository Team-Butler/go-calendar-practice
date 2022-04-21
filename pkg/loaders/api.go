package loaders

import (
	"go-calendar-practice/pkg/modules/example"

	"github.com/gin-gonic/gin"
)

func LoadApis(router *gin.Engine) {
	example.ExampleController(router)
}
