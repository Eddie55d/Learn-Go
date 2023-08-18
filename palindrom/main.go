package main

import "fmt"

func isPalindrom(b []byte) bool {
	for i := range b {
		if b[i] != b[(len(b)-1)-i] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	bytes := []byte{}
	for i, v := range str {
		bytes = append(bytes, byte(v))
		fmt.Printf("%v\n", bytes[i])
	}

	fmt.Printf("is palindrom: %v", isPalindrom(bytes))
}
