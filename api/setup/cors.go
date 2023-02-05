package setup

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	AllowedOrigins []string = []string{"*"}
	AllowedHeaders []string = []string{"Origin", "Authorization", "Content-Length", "Content-Type"}
	AllowedMethods []string = []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"}
)

func Cors(app *gin.RouterGroup) {
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     AllowedMethods,
		AllowHeaders:     AllowedHeaders,
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	app.Use(cors.New(config))
}
