package setup

import (
	"os"
	"path/filepath"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const defaultIndexHTMLContent = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Admin Panel</title>
  </head>
  <body>
    <h1>This should be the admin panel</h1>
  </body>
</html>`

func Static(app *gin.Engine) {
	path := filepath.Join(".", "static")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	path = filepath.Join(".", "static", "index.html")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		file.Write([]byte(defaultIndexHTMLContent))
		file.Close()
	}

	app.LoadHTMLGlob("static/*.html")
	app.Use(static.Serve("/static", static.LocalFile("static", false)))
}
