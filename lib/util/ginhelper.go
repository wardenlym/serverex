package util

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	isatty "github.com/mattn/go-isatty"
)

// GinLoggerMiddleware .
func GinLoggerMiddleware() gin.HandlerFunc {
	return loggerWithWriter(os.Stdout)
}

func loggerWithWriter(out io.Writer, notlogged ...string) gin.HandlerFunc {
	isTerm := true

	if w, ok := out.(*os.File); !ok ||
		(os.Getenv("TERM") == "dumb" || (!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd()))) {
		isTerm = false
	}

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			// Stop timer
			end := time.Now()
			latency := end.Sub(start)

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			// var statusColor, methodColor string
			if isTerm {
				// 	statusColor = colorForStatus(statusCode)
				// 	methodColor = colorForMethod(method)
			}
			comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

			fmt.Fprintf(out, "[GIN] %v | %3d | %13v | %15s |%s  %-7s \n%s",
				end.Format("2006/01/02 - 15:04:05"),
				statusCode,
				latency,
				clientIP,
				method,
				path,
				comment,
			)
		}
	}
}
