// Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
// и возвращающую количество полученных функцией аргументов и их сумму.

// Пример вызова вашей функции:

// a, b := sumInt(1, 0)
// fmt.Println(a, b)

// Результат: 2, 1

package main

import "fmt"

func sumInt(x ...int) (ln, sum int) {
	ln = len(x)

	for _, v := range x {
		sum += v
	}

	return ln, sum
}

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}
