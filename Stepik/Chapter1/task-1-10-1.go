// Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10.
//Квадрат каждого числа должен выводится в новой строке.

// Sample Input:

// Sample Output:

// 1
// 4
// 9
// 16
// 25
// 36
// 49
// 64
// 81
// 100

package main

import "fmt"

func main() {
	var num = 10

	for i := 1; i <= num; i++ {
		sum := i * i
		fmt.Println(sum)
	}
}
