package setup

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CustomLogger(app *gin.Engine) {
	logFileName := os.Getenv("API_LOG_FILE")
	if logFileName != "" {
		if strings.HasSuffix(os.Getenv("API_LOG_FILE"), ".log") {
			logFileName = os.Getenv("API_LOG_FILE")
		} else {
			logFileName = os.Getenv("API_LOG_FILE") + ".log"
		}
	} else {
		logFileName = "api.log"
	}

	logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file : %s\n", err)
		os.Exit(1)
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)

	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[API] %v | %3d | %13v | %15s | %-7s  %#v\n%s",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
}
