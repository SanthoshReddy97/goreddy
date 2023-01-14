package reddy

import (
	"github.com/SanthoshReddy97/goreddy/internal/httpReddy"
	"github.com/gin-gonic/gin"
)

func GinRouter() *gin.Engine {
	return httpReddy.GetGinRouter()
}

func RouterGroup() *gin.RouterGroup {
	return httpReddy.GetRouterGroup()
}
