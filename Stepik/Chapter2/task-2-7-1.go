// На вход подаются a и b - катеты прямоугольного треугольника.
// Нужно найти длину гипотенузы

// Sample Input:   6 8
// Sample Output:  10

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func hypotenuse(a, b float64) (float64, error) {

	if a == 0 || b == 0 {
		return 0, errors.New("zero is entered")
	}

	hypot := math.Sqrt(a*a + b*b)

	return hypot, nil

}

func main() {
	var x, y float64
	fmt.Scan(&x, &y)

	sum, err := hypotenuse(x, y)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
