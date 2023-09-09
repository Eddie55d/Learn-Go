package main

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

import (
	"crypto/sha256"
	"fmt"
)

type Entry struct {
	Key      string
	Value    string
	startIdx int
}

func strStr(haystack string, needle string) int {

	fmt.Println(haystack, needle)

	hArr := []byte(haystack)
	nArr := []byte(needle)

	H := sha256.New()

	H.Write([]byte(needle))

	needleHash := string(H.Sum(nil))

	needleArr := make([]*Entry, 0)

	endNeedle := 0

	for i := 0; i < len(hArr); i++ {
		H.Reset()
		if i+len(nArr) > len(hArr) {
			break
		}
		endNeedle = i + len(nArr)
		currentNeedle := hArr[i:endNeedle]
		H.Write(currentNeedle)
		currNeedleHash := string(H.Sum(nil))
		if currNeedleHash == needleHash {
			entryNeedle := &Entry{
				Key:      currNeedleHash,
				Value:    needleHash,
				startIdx: i,
			}
			needleArr = append(needleArr, entryNeedle)
		}
	}

	if len(needleArr) != 0 {
		return needleArr[0].startIdx
	}

	return -1
}

func main() {
	haystack1 := "sadbutsad"
	needle1 := "sad"

	haystack2 := "mississippi"
	needle2 := "issip"

	haystack3 := "mississippi"
	needle3 := "pi"

	haystack4 := "bbaa"
	needle4 := "aab"

	haystack5 := "babba"
	needle5 := "bbb"

	haystack6 := "abc"
	needle6 := "c"

	fmt.Println(strStr(haystack1, needle1)) // 0
	fmt.Println(strStr(haystack2, needle2)) // 4
	fmt.Println(strStr(haystack3, needle3)) // 9
	fmt.Println(strStr(haystack4, needle4)) // -1
	fmt.Println(strStr(haystack5, needle5)) // -1
	fmt.Println(strStr(haystack6, needle6)) // 2
}
