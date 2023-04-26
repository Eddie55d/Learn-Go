// Требуется написать программу, при выполнении которой
// с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B).
// Вывести сумму всех чисел от A до B  включительно.

// Sample Input: 1  5
// Sample Output: 15

package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	if a <= 100 && a < b {
		var sum = 0
		for a <= b {
			sum += a
			a++
		}
		fmt.Println(sum)
	} else {
		fmt.Println("incorrect input")
	}
}
