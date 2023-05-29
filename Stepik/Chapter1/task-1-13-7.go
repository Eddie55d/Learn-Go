// Количество нулей
// По данным числам, определите количество чисел, которые равны нулю.

// Входные данные
// Вводится натуральное число N, а затем N чисел.

// Выходные данные
// Выведите количество чисел, которые равны нулю.

// Sample Input:

// 5
// 1
// 8
// 100
// 0
// 12

// Sample Output:  1

package main

import "fmt"

func main() {

	var ln int
	fmt.Scan(&ln)

	a := make([]int, ln)
	sum := 0

	for i := 0; i < ln; i++ {
		fmt.Scan(&a[i])

		if a[i] == 0 {
			sum++
		}
	}
	fmt.Printf("%v", sum)

}
