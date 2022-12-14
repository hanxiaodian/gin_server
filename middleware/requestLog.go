package middleware

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger(logger logrus.FieldLogger, notLogged ...string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()

		c.Next()

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := logger.WithFields(logrus.Fields{
			"timestamp":     time.Now(),
			"requestMethod": c.Request.Method,
			"requestPath":   c.Request.URL.Path,
			"responseTime":  fmt.Sprintf("%d ms", latency),
			"statusCode":    statusCode,
			"responseSize":  fmt.Sprintf("%d B", dataLength),
			"clientIP":      clientIP,
			"referer":       referer,
			"userAgent":     clientUserAgent,
			"headers":       c.Request.Header,
			"qs":            c.Request.URL.Query(),
			"params":        c.Request.ParseForm(),
			"body":          c.Request.Body,
			// "resp-headers":  c.Request.Response.Header,
			// "return": c.Request.Response.Body,
		})

		//if len(c.Errors) > 0 {
		if c.Errors != nil {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%dms", latency)
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
