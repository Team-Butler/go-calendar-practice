package loaders

import (
	"go-calendar-practice/pkg/modules/base"
	"go-calendar-practice/pkg/modules/example"

	"github.com/gin-gonic/gin"
)

func LoadApis(router *gin.Engine) {
	base.AddController(router)

	example.AddController(router)
}
