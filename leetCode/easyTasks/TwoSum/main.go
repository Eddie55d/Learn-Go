package main

// https://leetcode.com/problems/two-sum/

import "fmt"

func twoSum(nums []int, target int) []int {
	idx := []int{}

	for i := range nums {
		for x := range nums {
			if nums[i]+nums[x] == target && i != x {
				idx = append(idx, i)
				idx = append(idx, x)
				return idx
			}
		}
	}
	return idx
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	nums1 := []int{3, 2, 4}
	target1 := 6

	fmt.Println(twoSum(nums, target))   // [0, 1]
	fmt.Println(twoSum(nums1, target1)) // [1, 2]
}
