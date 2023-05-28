// Последовательность Фибоначчи определена
// следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.

// Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...

// Напишите функцию, которая по указанному натуральному n возвращает φn.

// Входные данные

// Вводится одно число n.

// Выходные данные

// Необходимо вывести  значение φn.

// Sample Input:  4

// Sample Output:  3

package main

import (
	"fmt"
)

func fibonacci(n int) int {

	fibbo := []int{1, 1}

	prevFib := 1
	fib := 1
	nextFib := 0

	for i := 2; i <= n; i++ {
		nextFib = prevFib + fib
		fibbo = append(fibbo, nextFib)
		prevFib, fib = fib, nextFib
	}

	return fibbo[n-1]

}

func main() {
	var num int

	fmt.Scan(&num)

	if num < 1 {
		return
	}

	fmt.Println(fibonacci(num))
}
