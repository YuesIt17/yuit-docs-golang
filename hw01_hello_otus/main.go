// Package main печатает перевёрнутую фразу "Hello, OTUS!".
package main

import (
	"fmt"

	//nolint:depguard // По условию ДЗ используем пакет reverse.
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello, OTUS!"))
}
