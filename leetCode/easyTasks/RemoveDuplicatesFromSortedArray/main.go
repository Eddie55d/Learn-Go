package main

// https://leetcode.com/problems/merge-two-sorted-lists/

import (
	"fmt"
	"sort"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	fmt.Println(nums)

	arr := []int{}
	var isRepeat bool

	for i, v := range nums {
		current := v
		isRepeat = false
		if len(arr) == 0 {
			arr = append(arr, v)
		} else {

			for _, val := range arr {
				if current == val {
					isRepeat = true
					nums[i] = 777
					break
				}
			}
			if !isRepeat {
				arr = append(arr, current)
			}

		}

	}

	fmt.Println(nums)
	fmt.Println(arr)
	k := len(arr)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums)

	return k
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(nums)) // 5
}
