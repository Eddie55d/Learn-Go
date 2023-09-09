package main

// https://leetcode.com/problems/length-of-last-word/

import "fmt"

func lengthOfLastWord(s string) int {
	strArr := []byte(s)

	space := byte(' ')

	endStr := len(strArr) - 1

	lastWord := make([]byte, 0)

	var startWord bool

	for i := endStr; i >= 0; i-- {
		if startWord && strArr[i] == space {
			break
		}
		if strArr[i] != space {
			startWord = true
			lastWord = append(lastWord, strArr[i])
		}
	}

	if len(lastWord) != 0 {
		return len(lastWord)
	}

	return 0
}

func main() {
	s := "Hello World"
	s1 := "   fly me   to   the moon  "
	s2 := "a"

	fmt.Println(lengthOfLastWord(s))  // 5
	fmt.Println(lengthOfLastWord(s1)) // 4
	fmt.Println(lengthOfLastWord(s2)) // 1
}
