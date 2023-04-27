// Дано трехзначное число. Переверните его, а затем выведите.

// Формат входных данных
// На вход дается трехзначное число, не оканчивающееся на ноль.

// Формат выходных данных
// Выведите перевернутое число.

// Sample Input:  843
// Sample Output:  348

package main

import "fmt"

func main() {

	var num int
	fmt.Scan(&num)

	if num < 1 && num > 100 {
		return
	}

	a1 := num / 100
	a2 := (num / 10) % 10
	a3 := num % 10

	fmt.Printf("%d%d%d", a3, a2, a1)
}
