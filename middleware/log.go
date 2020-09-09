package middleware

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/rifflock/lfshook"

	"github.com/gin-gonic/gin"
	rotalog "github.com/lestrrat-go/file-rotatelogs"
	_ "github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	filePath := "log/blog.log"
	linkName := "latest_log.log"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := rotalog.New(
		filePath+"%Y%m%d.log",
		rotalog.WithMaxAge(7*24*time.Hour),
		rotalog.WithRotationTime(7*24*time.Hour),
		rotalog.WithLinkName(linkName),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}

		method := c.Request.Method
		path := c.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"host":      hostName,
			"status":    statusCode,
			"spendtime": spendTime,
			"IP":        clientIp,
			"method":    method,
			"url":       path,
			"datasize":  dataSize,
			"agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
