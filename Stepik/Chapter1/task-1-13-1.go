// Дано трехзначное число. Найдите сумму его цифр.

// Формат входных данных
// На вход дается трехзначное число.

// Формат выходных данных
// Выведите одно целое число - сумму цифр введенного числа.

// Sample Input:  745

// Sample Output:  16

package main

import "fmt"

func main() {

	var num int
	fmt.Scan(&num)

	if num < 1 || num > 100 {
		return
	}

	a1 := num / 100
	a2 := (num / 10) % 10
	a3 := num % 10

	sum := a1 + a2 + a3

	fmt.Println(sum)
}
