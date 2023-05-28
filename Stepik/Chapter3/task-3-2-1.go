// Напишите функцию с именем convert, которая конвертирует входное
// число типа int64 в uint16.

// Считывать и выводить ничего не нужно!

// Считайте что функция main и пакет main уже объявлены!

package main

import "fmt"

func convert(a int64) uint16 {
	result := uint16(a)
	return result
}

func main() {
	var num int64
	fmt.Scan(&num)

	a := convert(num)

	fmt.Printf("%T", a)
}
