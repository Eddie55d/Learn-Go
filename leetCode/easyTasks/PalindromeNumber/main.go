package main

// https://leetcode.com/problems/palindrome-number/

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	arr := []byte(strX)
	for i := range arr {
		if arr[i] != arr[(len(arr)-1)-i] {
			return false
		}
	}
	return true
}

func main() {
	x := 121
	x1 := -121

	fmt.Println(isPalindrome(x))  // true
	fmt.Println(isPalindrome(x1)) // false
}
