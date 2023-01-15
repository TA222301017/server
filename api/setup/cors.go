package setup

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	AllowedOrigins []string = []string{"*"}
	AllowedHeaders []string = []string{"Authorization"}
	AllowedMethods []string = []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"}
)

func Cors(app *gin.Engine) {
	config := cors.Config{
		AllowOrigins:     AllowedOrigins,
		AllowMethods:     AllowedMethods,
		AllowHeaders:     AllowedHeaders,
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}
	app.Use(cors.New(config))
}
