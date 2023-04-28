// Напишите "функцию голосования", возвращающую то значение (0 или 1),
// которое среди значений ее аргументов x, y, z встречается чаще.

// Входные данные

// Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).

// Выходные данные

// Необходимо вернуть значение функции от x, y и z.

// Sample Input:   0 0 1

// Sample Output:  0

package main

import "fmt"

func vote(x, y, z int) int {
	count1 := 0
	count2 := 0

	nums := [3]int{x, y, z}

	for _, v := range nums {
		if v == 0 {
			count1++
		}
	}

	for _, v := range nums {
		if v == 1 {
			count2++
		}
	}

	total := 0

	switch {
	case count1 > count2:
		total = 0
	case count2 > count1:
		total = 1
	}

	return total
}

func main() {
	fmt.Println(vote(1, 1, 0))
	fmt.Println(vote(0, 0, 1))
}
