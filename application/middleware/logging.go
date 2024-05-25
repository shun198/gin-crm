package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleWare() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] IP: %s %s %s %d 実行時間: %s \n",
			params.TimeStamp.Format(time.RFC3339),
			params.ClientIP,
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}
