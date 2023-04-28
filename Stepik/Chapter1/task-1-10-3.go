// Напишите программу, которая в последовательности чисел находит сумму двузначных чисел,
// кратных 8. Программа в первой строке получает на вход число n - количество чисел
// в последовательности, во второй строке -- n чисел, входящих в данную последовательность.

// Sample Input:
// 5
// 38 24 800 8 16

// Sample Output: 40

package main

import "fmt"

func main() {
	var lenNums, nums int

	fmt.Scan(&lenNums)

	sum := 0

	for i := 1; i <= lenNums; i++ {
		fmt.Scan(&nums)
		if nums > 9 && nums < 100 && nums%8 == 0 {
			sum += nums
		}
	}
	fmt.Println(sum)
}
