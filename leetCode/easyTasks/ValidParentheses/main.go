package main

// https://leetcode.com/problems/valid-parentheses/

import "fmt"

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	fmt.Println(s)

	ascii1 := byte('(')
	ascii2 := byte(')')
	ascii3 := byte('[')
	ascii4 := byte(']')
	ascii5 := byte('{')
	ascii6 := byte('}')

	byteArr := []byte{}

	byteStr := []byte(s)

	for _, v := range byteStr {
		switch {
		case v == ascii1 || v == ascii3 || v == ascii5:
			byteArr = append(byteArr, v)
		case v == ascii2 && len(byteArr) == 0 || v == ascii4 && len(byteArr) == 0 || v == ascii6 && len(byteArr) == 0:
			return false
		case v == ascii2:
			if byteArr[len(byteArr)-1] == ascii1 {
				byteArr[len(byteArr)-1] = 0

				byteArr = byteArr[:len(byteArr)-1]
			} else {
				return false
			}

		case v == ascii4:
			if byteArr[len(byteArr)-1] == ascii3 {
				byteArr[len(byteArr)-1] = 0
				byteArr = byteArr[:len(byteArr)-1]
			} else {
				return false
			}
		case v == ascii6:
			if byteArr[len(byteArr)-1] == ascii5 {
				byteArr[len(byteArr)-1] = 0

				byteArr = byteArr[:len(byteArr)-1]
			} else {
				return false
			}
		}
	}

	if len(byteArr) != 0 {
		return false
	}

	return true
}

func main() {
	s1 := "()"
	s2 := "()[]{}"
	s3 := "(]"
	s4 := "{[]}"
	s5 := "(){}}{"
	s6 := "([])"
	s7 := "{()}"
	s8 := "[()]"
	s9 := "[{}]"
	s10 := "({})"
	s11 := "([)]"

	fmt.Println(isValid(s1))  // true
	fmt.Println(isValid(s2))  // true
	fmt.Println(isValid(s3))  // false
	fmt.Println(isValid(s4))  // true
	fmt.Println(isValid(s5))  // false
	fmt.Println(isValid(s6))  // true
	fmt.Println(isValid(s7))  // true
	fmt.Println(isValid(s8))  // true
	fmt.Println(isValid(s9))  // true
	fmt.Println(isValid(s10)) // true
	fmt.Println(isValid(s11)) // false
}
