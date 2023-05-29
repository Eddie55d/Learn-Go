// По данному числу N распечатайте все целые значения степени двойки,
// не превосходящие N, в порядке возрастания.

// Входные данные

// Вводится натуральное число.

// Выходные данные

// Выведите ответ на задачу.

// Sample Input:  50

// Sample Output:  1 2 4 8 16 32

package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Scan(&num)

	var lastNum float64

	for i := 0; lastNum < float64(num); i++ {
		pow := math.Pow(2, float64(i))

		lastNum = pow

		if lastNum <= float64(num) {
			fmt.Printf("%d ", int(lastNum))
		}

	}
}
