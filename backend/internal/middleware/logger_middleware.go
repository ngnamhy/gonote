package middleware

import (
	"bytes"
	mylog "gonote/pkg/log"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(data []byte) (n int, err error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

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

		customWriter := &CustomResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}

		c.Writer = customWriter

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		if raw != "" {
			path = path + "?" + raw
		}

		event := mylog.HTTPLogger.Info()

		responseBody := customWriter.body.Bytes()
		var responseBodyStr string
		if len(responseBody) == 0 {
			responseBodyStr = "[empty response]"
		} else if len(responseBody) > maxBodySize {
			responseBodyStr = "[response too large to log]"
		} else {
			responseBodyStr = string(responseBody)
		}

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
			Dur("latency", latency).
			Str("response_body", responseBodyStr)

		if len(bodyBytes) > 0 && string(bodyBytes) != "[body too large to log]" {
			bodyStr := string(bodyBytes)
			logEvent = logEvent.Str("request_body", bodyStr)
		} else if len(bodyBytes) > 0 {
			logEvent = logEvent.Str("request_body", string(bodyBytes))
		}

		logEvent.Msg("HTTP request")
	}
}
