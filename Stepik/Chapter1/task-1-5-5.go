// Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).

// Формат входных данных
// На вход дается натуральное число, не превосходящее 10000.

// Формат выходных данных
// Выведите одно целое число - число десятков.

// Sample Input: 2010
// Sample Output: 1

package main

import (
	"fmt"
	"math"
)

var a int

func main() {
	fmt.Scan(&a)
	b := float64(a) / 100
	//fmt.Println(b)
	_, c := math.Modf(b)
	//fmt.Printf("%.1f\n", c)
	d := float64(c) * 10
	e := int(d)
	fmt.Println(e)
}
