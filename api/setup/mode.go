package setup

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Mode() {
	mode := os.Getenv("APP_MODE")
	if mode == "dev" || mode == "" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
