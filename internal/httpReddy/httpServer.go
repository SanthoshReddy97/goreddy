package httpReddy

import (
	"github.com/gin-gonic/gin"
)

var v1 = &gin.RouterGroup{}
var router = &gin.Engine{}

func StartServer() {
	router = RouterSetUp()
	v1 = router.Group("")
	router.Run()
}

func RouterSetUp() *gin.Engine {
	router := gin.Default()
	return router
}

func GetGinRouter() *gin.Engine {
	return router
}

func GetRouterGroup() *gin.RouterGroup {
	return v1
}
