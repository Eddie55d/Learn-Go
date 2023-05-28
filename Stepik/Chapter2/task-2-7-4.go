// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и
// вывести получившееся число.

// Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1.
// Единица в квадрате - 1. В итоге получаем 811181

// Sample Input:   9119

// Sample Output:  811181

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	var num int

	fmt.Scan(&num)

	str := strconv.Itoa(num)

	lenStr := utf8.RuneCountInString(str)

	if lenStr > 1000 {
		panic("строка больше 1000 символов")
	}

	formatStr := []string{}

	for _, v := range str {
		formatStr = append(formatStr, string(v))
	}

	intStr := []string{}

	for _, v := range formatStr {
		s, _ := strconv.Atoi(v)
		d := strconv.Itoa(s * s)
		intStr = append(intStr, d)
	}

	total := strings.Join(intStr, "")

	fmt.Println(total)

}
