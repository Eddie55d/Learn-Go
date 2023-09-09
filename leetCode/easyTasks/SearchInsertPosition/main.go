package main

// https://leetcode.com/problems/search-insert-position/

import "fmt"

func searchInsert(nums []int, target int) int {

	low := 0

	high := len(nums) - 1

	if len(nums) == 0 {
		return 0
	}

	if target == nums[0] {
		return 0
	}

	if target == nums[high] {
		return high
	}

	for low <= high {
		mid := (low + high) >> 1
		midVal := nums[mid]

		if high-low <= 1 && midVal != target {

			if nums[low] == target {
				return low
			} else if nums[high] == target {
				return high
			} else if nums[high] < target {
				return high + 1
			} else if nums[low] > target && low < mid {
				return low - 1
			} else if target > nums[mid] && low == mid {
				return mid + 1
			} else if target < nums[mid] && low == mid {
				return mid
			}

		} else if midVal < target {
			low = mid + 1

		} else if midVal == target {
			return mid

		} else {
			high = mid - 1
		}
	}

	return 0
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5

	nums1 := []int{1, 3, 5, 6}
	target1 := 2

	nums2 := []int{1, 3}
	target2 := 2

	nums3 := []int{3, 5, 7, 9, 10}
	target3 := 8

	fmt.Println(searchInsert(nums, target))   // 2
	fmt.Println(searchInsert(nums1, target1)) // 1
	fmt.Println(searchInsert(nums2, target2)) // 1
	fmt.Println(searchInsert(nums3, target3)) // 3
}
