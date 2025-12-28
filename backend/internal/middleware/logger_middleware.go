package middleware

import (
	"bytes"
	mylog "gonote/pkg/log"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

const maxBodySize = 10 * 1024

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		method := c.Request.Method

		var bodyBytes []byte
		if c.Request.Body != nil {
			if c.Request.ContentLength > 0 && c.Request.ContentLength <= maxBodySize {
				bodyBytes, _ = io.ReadAll(c.Request.Body)
			} else if c.Request.ContentLength > maxBodySize {
				bodyBytes = []byte("[body too large to log]")
			}
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		if raw != "" {
			path = path + "?" + raw
		}

		event := mylog.HTTPLogger.Info()
		if len(c.Errors) > 0 || statusCode >= 500 {
			event = mylog.HTTPLogger.Error()
		} else if statusCode >= 400 {
			event = mylog.HTTPLogger.Warn()
		}

		logEvent := event.
			Str("method", method).
			Str("path", path).
			Int("status", statusCode).
			Str("ip", clientIP).
			Str("user_agent", userAgent).
			Dur("latency", latency)

		if len(bodyBytes) > 0 && string(bodyBytes) != "[body too large to log]" {
			bodyStr := string(bodyBytes)
			logEvent = logEvent.Str("body", bodyStr)
		} else if len(bodyBytes) > 0 {
			logEvent = logEvent.Str("body", string(bodyBytes))
		}

		logEvent.Msg("HTTP request")
	}
}
