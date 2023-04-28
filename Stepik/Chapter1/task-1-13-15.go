// Из натурального числа удалить заданную цифру.

// Входные данные

// Вводятся натуральное число и цифра, которую нужно удалить.

// Выходные данные

// Вывести число без заданных цифр.

// Sample Input:

// 38012732
// 3

// Sample Output:  801272

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	c := strconv.Itoa(a)

	d := []int{}

	for _, v := range c {
		x, _ := strconv.Atoi(string(v))
		d = append(d, x)
	}

	for _, v := range d {
		if v != b {
			fmt.Printf("%v", v)
		}
	}
}
