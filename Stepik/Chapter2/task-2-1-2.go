// Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

// Входные данные

// Вводится четыре числа.

// Выходные данные

// Необходимо вернуть из функции наименьшее из 4-х данных чисел

// Sample Input:  4 5 6 7

// Sample Output:  4

package main

import "fmt"

func minimumFromFour() int {
	nums := []int{}
	var n int

	for i := 0; i < 4; i++ {
		fmt.Scan(&n)
		nums = append(nums, n)
	}

	result := nums[0]
	for _, elem := range nums {
		if elem < result {
			result = elem
		}
	}

	return result
}

func main() {
	//min := minimumFromFour()
	fmt.Println(minimumFromFour())
}
