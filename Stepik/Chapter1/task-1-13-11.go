// По данному числу n закончите фразу "На лугу пасется..." одним из возможных
// продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

// Входные данные

// Дано число n (0<n<100).

// Выходные данные

// Программа должна вывести введенное число n и одно
// из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov.
// Между числом и словом должен стоять ровно один пробел.

// Sample Input:   10
// Sample Output:  10 korov

package main

import "fmt"

func main() {

	var num int
	fmt.Scan(&num)

	if num > 99 && num < 1 {
		return
	}

	a := "korova"
	b := "korovy"
	c := "korov"

	switch {

	case num%10 == 1 && num != 11 || num == 1:
		fmt.Printf("%d %s", num, a)

	case num > 4 && num < 21:
		fmt.Printf("%d %s", num, c)

	case num <= 2 && num > 5 || num%10 > 1 && num%10 < 5:
		fmt.Printf("%d %s", num, b)

	default:
		fmt.Printf("%d %s", num, c)
	}

}
