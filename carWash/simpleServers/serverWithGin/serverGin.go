package serverGin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var message string = "Hello, World! I am Gin!"

func ServerGin() {
	r := gin.Default()
	r.GET("/hello-world", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", message)
	})
	r.Run("localhost:8082") // listen and serve on 0.0.0.0:8082 (for windows "localhost:8082")
}
