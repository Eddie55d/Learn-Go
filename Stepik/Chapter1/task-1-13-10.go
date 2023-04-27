// Самое большое число, кратное 7
// Найдите самое большее число на отрезке от a до b, кратное 7 .

// Входные данные
// Вводится два целых числа a и b (a≤b).

// Выходные данные
// Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b),
// кратное 7 , или выведите "NO" - если таковых нет.

// Sample Input:

// 100
// 500

// Sample Output:  497

package main

import "fmt"

func main() {

	var a, b int

	fmt.Scan(&a, &b)

	if a > b {
		return
	}

	c := []int{}

	for i := a; i <= b; i++ {
		if i%7 == 0 {
			c = append(c, i)
		}
	}

	switch {
	case len(c) == 0:
		fmt.Println("NO")

	case len(c) > 1:
		end := int(c[len(c)-1])
		fmt.Printf("%v", end)

	case len(c) == 1:
		fmt.Printf("%v", c[0])
	}
}
