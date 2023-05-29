// Номер числа Фибоначчи
// Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является,
// то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи,
// выведите число -1.

// Входные данные

// Вводится натуральное число.

// Выходные данные

// Выведите ответ на задачу.

// Sample Input:  8

// Sample Output:  6

package main

import (
	"fmt"
)

func main() {
	var A int
	fmt.Scan(&A)

	if A < 1 {
		return
	}

	fibbo := []int{0: 0, 1: 1}

	prevFib := 0
	fib := 1
	nextFib := 0

	for i := 0; nextFib <= A; i++ {
		nextFib = prevFib + fib
		fibbo = append(fibbo, nextFib)
		prevFib = fib
		fib = nextFib
	}

	idx := 0

	for i, v := range fibbo {
		if v == A {
			idx = i
			fmt.Printf("%v", idx)
		}
	}

	if idx == 0 {
		fmt.Println("-1")
	}
}
