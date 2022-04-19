package loaders

import (
	"go-calendar-practice/pkg/modules/example"

	"github.com/gin-gonic/gin"
)

func LoadApis(r *gin.Engine) {
	example.ExampleController(r)
}