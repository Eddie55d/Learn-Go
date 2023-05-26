//запустить main и менять только порт
// 8080: "Hello, World! I am Go!"
// 8081: "Hello, World! I am Gorilla!"
// 8082: "Hello, World! I am Gin!"
// 8083: "Hello, World! I am Fast!"

package main

import (
	"example.com/serverFast"
	"example.com/serverGin"
	"example.com/serverGo"
	"example.com/serverGorilla"
)

func main() {
	func() {
		go serverGorilla.ServerGorilla()
	}()

	func() {
		go serverGo.ServerGO()
	}()

	func() {
		go serverFast.ServerFast()
	}()

	serverGin.ServerGin()
}
