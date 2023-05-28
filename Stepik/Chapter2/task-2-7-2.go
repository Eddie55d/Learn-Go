// Дана строка, содержащая только английские буквы (большие и маленькие).
// Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после
// 	последней символ ‘*’ добавлять не нужно).

// Входные данные

// Вводится строка ненулевой длины. Известно также, что длина строки не
// превышает 1000 знаков.

// Выходные данные

// Вывести строку, которая получится после добавления символов '*'.

// Sample Input:

// LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO

// Sample Output:

// L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O

package main

import (
	"fmt"
	"strings"
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

	fmt.Println(strings.Join(formatStr, "*"))
}
