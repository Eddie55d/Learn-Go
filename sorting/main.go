package main

import "fmt"

// bubble sort
func bubbleSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {

		var swap bool

		for idx := 0; idx < len(arr)-i-1; idx++ {

			current := arr[idx]
			next := arr[idx+1]

			if current > next {
				swap = true
				arr[idx] = next
				arr[idx+1] = current

			}
		}

		if !swap {
			break
		}

	}

	return arr
}

// insert sort
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}

	return arr
}

// merge sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	sortArr := make([]int, 0, len(left)+len(right))

	leftIdx := 0
	rightIdx := 0

	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] < right[rightIdx] {
			sortArr = append(sortArr, left[leftIdx])
			leftIdx++
		} else {
			sortArr = append(sortArr, right[rightIdx])
			rightIdx++
		}
	}

	sortArr = append(sortArr, left[leftIdx:]...)
	sortArr = append(sortArr, right[rightIdx:]...)

	return sortArr
}

// quick sort
func quickSort(arr []int, lo, hi int) []int {

	if lo >= hi {
		return arr
	}

	p := partition(arr, lo, hi)

	quickSort(arr, lo, p-1)
	quickSort(arr, p, hi)

	return arr
}

func partition(arr []int, lo, hi int) int {
	mid := arr[(lo+hi)/2]
	i := lo
	j := hi

	for i < j {
		for arr[i] <= mid && i != j {
			i++
			if arr[i] >= mid {
				break
			}
		}
		for arr[j] > mid && i != j {
			j--
			if arr[j] <= mid {
				break
			}
		}

		if i >= j {
			return j
		}

		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp

	}

	return j
}

func main() {
	arr := []int{5, 4, 2, 1, 8, 9, 1, 6, 3, 7}
	arr2 := []int{5, 4, 2, 1, 8, 9, 1, 6, 3, 7}
	arr3 := []int{5, 4, 2, 1, 8, 9, 1, 6, 3, 7}
	arr4 := []int{5, 4, 2, 1, 8, 9, 1, 6, 3, 7}

	fmt.Println("source array:", arr)
	fmt.Println("bubbleSort array:", bubbleSort(arr))
	fmt.Println("insertionSort array:", insertionSort(arr2))
	fmt.Println("mergeSort array:", mergeSort(arr3))
	fmt.Println("quickSort array:", quickSort(arr4, 0, len(arr4)-1))
}
