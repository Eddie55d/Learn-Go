// Количество минимумов
// Найдите количество минимальных элементов в последовательности.

// Входные данные

// Вводится натуральное число N, а затем N целых чисел последовательности.

// Выходные данные

// Выведите количество минимальных элементов последовательности.

// Sample Input:

// 3
// 21
// 11
// 4

// Sample Output: 1

package main

import "fmt"

func main() {
	var ln, a int

	fmt.Scan(&ln)
	array := make([]int, ln)

	for i := 0; i < ln; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	sum := 0
	minEl := array[0]

	for _, elem := range array {
		if elem < minEl {
			minEl = elem
			sum = 1
		} else if elem == minEl {
			sum++
		}
	}

	fmt.Printf("%v", sum)
}
