// Напишите функцию, которая умножает значения на которые ссылаются два указателя
//  и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.

// Sample Input:  2 2

// Sample Output: 4

package main

import "fmt"

func test(x1 *int, x2 *int) {
	fmt.Printf("%d", *x1**x2)

}

func main() {
	x1 := 2
	x2 := 2

	test(&x1, &x2)
}
