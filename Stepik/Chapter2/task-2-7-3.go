// Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

// Входные данные

// Вводится строка ненулевой длины. Известно также, что длина строки
// не превышает 1000 знаков и строка содержит только арабские цифры.

// Выходные данные

// Выведите максимальную цифру, которая встречается во введенной строке.

// Sample Input:    1112221112
// Sample Output:   2

package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var str string

	fmt.Scan(&str)

	lenStr := utf8.RuneCountInString(str)

	if lenStr > 1000 {
		panic("строка больше 1000 символов")
	}

	formatStr := []string{}

	for _, v := range str {
		formatStr = append(formatStr, string(v))
	}

	intStr := []int{}

	for _, v := range formatStr {
		s, _ := strconv.Atoi(v)
		intStr = append(intStr, s)
	}

	prev := intStr[0]
	for _, v := range intStr {
		if v > prev {
			prev = v
		}
	}

	fmt.Println(prev)

}
