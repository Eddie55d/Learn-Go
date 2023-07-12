package logger

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggerInfo(c *gin.Context) {
	t := time.Now()
	reqMethod := c.Request.Method
	clientIP := c.ClientIP()
	path := c.Request.URL.Path

	c.Next()

	statusCode := strconv.Itoa(c.Writer.Status())
	respTime := time.Since(t)

	log.WithFields(log.Fields{
		"method":        reqMethod,
		"client-ip":     clientIP,
		"path":          path,
		"statusCode":    statusCode,
		"response-time": respTime.String(),
	}).Info()

}
