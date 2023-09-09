package main

// https://leetcode.com/problems/remove-element/

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	arr := []int{}

	for _, v := range nums {
		current := v

		if current != val {
			arr = append(arr, v)
		}

	}
	fmt.Println(nums)
	fmt.Println(arr)

	k := len(arr)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i, v := range arr {
		nums[i] = v
	}

	if len(nums) != len(arr) {
		for i := len(arr); i < len(nums); i++ {
			nums[i] = val
		}
	}

	fmt.Println(nums)

	return k

}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}

	val := 2

	fmt.Println(removeElement(nums, val)) // 5
}
