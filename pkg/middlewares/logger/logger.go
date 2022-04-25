package logger

import (
	"fmt"
	"time"

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
	var statusColor, methodColor, resetColor string

	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	return fmt.Sprintf("%s[%s]%s %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		cyan, serviceName, resetColor,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
