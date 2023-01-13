package templates

func Main() []byte {
	data := []byte(`package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`)
	return data
}

func Mod() []byte {
	data := []byte(``)
	return data
}
