package logger

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type LogFormatter func(params gin.LogFormatterParams) string

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return logFormatter(param)
	})
}

var logFormatter = func(param gin.LogFormatterParams) string {
	const serviceName = "TASKAT"
	var serviceColor, statusColor, methodColor, resetColor string

	if param.IsOutputColor() {
		serviceColor = cyan
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	return fmt.Sprintf("%s[%s]%s %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		serviceColor, serviceName, resetColor,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
