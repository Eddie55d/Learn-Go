// Задача 1: Разворот массива
// Напишите функцию на Go, которая принимает на вход массив целых чисел и возвращает новый массив с обратным порядком элементов.

// Пример:

package main

import "fmt"

func reverse(arr []int) []int {
	count := len(arr)
	var arr2 []int
	for i := count - 1; i >= 0; i-- {
		arr2 = append(arr2, arr[i])

	}

	return arr2
}

func main() {
	arr := []int{10, 20, 30, 40}
	fmt.Println(reverse(arr))
}

// Ожидаемый вывод:

// [5 4 3 2 1]
