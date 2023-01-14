package templates

func Main() []byte {
	data := []byte(`package main

import (
	"net/http"
  
	"github.com/santhoshreddy/goreddy"
)

func main() {
	goreddy.gin.ExecuteFromCommandLine()
}
`)
	return data
}

func Mod() []byte {
	data := []byte(``)
	return data
}
