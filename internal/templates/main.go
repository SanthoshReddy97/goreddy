package templates

func Main() []byte {
	data := []byte(`package main

import (
	"github.com/SanthoshReddy97/goreddy/src/reddy"
)

func main() {
	reddy.ExecuteFromCommandLine()
}
`)
	return data
}

func Settings(appName string) []byte {
	data := []byte(`package ` + appName + `

const (
	SECRET_KEY = "b8=1_&3t4#fojvav69*_@f2migm10grc2$n770z5zjy8tkuu0g"
)

var ALLOWED_HOSTS = []string{"*"}
	`)
	return data
}

func Urls(appName string) []byte {
	data := []byte(`package ` + appName + `

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
	`)
	return data
}
