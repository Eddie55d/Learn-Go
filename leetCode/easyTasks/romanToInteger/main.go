package main

// https://leetcode.com/problems/roman-to-integer/

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {

	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	val := strings.Split(s, "")

	next := 0
	sum := 0
	var jump bool
	for i := 0; i <= len(val)-1; i++ {

		for key, value := range roman {

			v := val[i]

			if i != len(val)-1 {
				n := val[i+1]
				next = roman[n]
			} else {
				next = 0
			}

			if v == key {

				if next <= value {
					sum += value
					jump = false
					break
				}
				if next > value {
					q := next - value
					sum += q
					jump = true
					break
				}
			}
		}
		if jump {
			i++
		}
	}
	return sum
}

func main() {

	s := "MCMXCIV"
	s2 := "III"
	s3 := "LVIII"

	fmt.Println(romanToInt(s))  // 1994
	fmt.Println(romanToInt(s2)) // 3
	fmt.Println(romanToInt(s3)) // 58
}
