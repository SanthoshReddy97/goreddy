package templates

import (
	"net/http"

	"github.com/SanthoshReddy97/goreddy/src/reddy"
	"github.com/gin-gonic/gin"
)

func Router() {
	reddy.GinRouter().GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
}
