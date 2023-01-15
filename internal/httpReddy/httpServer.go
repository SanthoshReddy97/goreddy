package httpReddy

import (
	"github.com/gin-gonic/gin"
)

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*gin.Context)

type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}

var URLs []RouteInfo

func StartServer() {
	router := gin.Default()
	for _, url := range URLs {
		if url.Method == "GET" {
			router.GET(url.Path, gin.HandlerFunc(url.HandlerFunc))
		}
	}
	router.Run()
}
