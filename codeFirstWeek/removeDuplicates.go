// Задача 2: Удаление дубликатов из массива
// Напишите функцию на Go, которая принимает на вход массив целых чисел и возвращает новый массив,
// содержащий только уникальные значения из исходного массива.
// Порядок элементов в новом массиве должен сохраняться.

// Пример:

package main

import "fmt"

func removeDuplicates(arr []int) []int {
	//num:= arr[0]
	var arr2 = []int{}

	for _, v := range arr {
		var isDupl bool
		for _, z := range arr2 {
			if v == z {
				isDupl = true
				break
			}
		}
		if !isDupl {
			arr2 = append(arr2, v)
		}

	}

	return arr2
}

func main() {
	arr := []int{1, 2, 3, 2, 1, 5, 6, 5, 7}
	fmt.Println(removeDuplicates(arr))
}

//Ожидаемый вывод:

//[1 2 3 5 6 7]
