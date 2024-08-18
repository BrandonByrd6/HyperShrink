package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		requestMethod := ctx.Request.Method

		requestedUrlPath := ctx.Request.URL.Path

		statusCode := ctx.Writer.Status()

		clientIp := ctx.ClientIP()

		log.WithFields(log.Fields{
			"METHOD":    requestMethod,
			"URI":       requestedUrlPath,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIp,
		}).Info("Http Request")
	}
}
