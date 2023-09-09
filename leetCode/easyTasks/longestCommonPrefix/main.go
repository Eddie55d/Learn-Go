package main

// https://leetcode.com/problems/longest-common-prefix/

import (
	"fmt"
)

type Pref struct {
	value string
	idx   int
	count int
}

func longestCommonPrefix(strs []string) string {

	for i := range strs {
		if len(strs[i]) == 0 {
			return ""
		}
	}

	if len(strs) == 1 {
		return strs[0]
	}

	lenStr := len(strs[0])

	for _, x := range strs {
		currentLen := len(x)
		if currentLen <= lenStr {
			lenStr = currentLen
		}
	}

	prefArr := []Pref{}

	for i := 0; i < lenStr; i++ {
		idx := strs[0][i]
		for _, val := range strs {
			if idx == val[i] {
				prefElem := &Pref{
					value: string(idx),
					idx:   i,
					count: 1,
				}
				if len(prefArr) != 0 {
					var flag bool
					for prefIdxArr, elem := range prefArr {
						if elem.value == prefElem.value && elem.idx == prefElem.idx {
							prefElem.count = elem.count + 1
							flag = true
							prefArr[prefIdxArr] = *prefElem
							break
						}
					}
					if !flag {
						prefArr = append(prefArr, *prefElem)
					}
				} else {
					prefArr = append(prefArr, *prefElem)
				}
			}
		}
	}

	prefResult := ""

	for _, v := range prefArr {
		if v.count == len(strs) {
			prefResult += v.value
		} else {
			break
		}
	}

	return prefResult
}
func main() {

	strs := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}
	strs3 := []string{"dog"}
	strs4 := []string{"cir", "car"}
	strs5 := []string{"aa", "aa"}
	strs6 := []string{"caa", "", "a", "acb"}
	strs7 := []string{"c", "acc", "ccc"}
	strs8 := []string{"ab", "caca", "accc", "cbcb"}

	fmt.Println(longestCommonPrefix(strs))  //fl
	fmt.Println(longestCommonPrefix(strs2)) // ""
	fmt.Println(longestCommonPrefix(strs3)) // dog
	fmt.Println(longestCommonPrefix(strs4)) // c
	fmt.Println(longestCommonPrefix(strs5)) // aa
	fmt.Println(longestCommonPrefix(strs6)) // ""
	fmt.Println(longestCommonPrefix(strs7)) // ""
	fmt.Println(longestCommonPrefix(strs8)) // ""
}
