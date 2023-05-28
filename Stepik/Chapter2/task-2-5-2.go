// На вход подается строка. Нужно определить, является ли она палиндромом.
// Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
// Все входные строчки в нижнем регистре.

// Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих
// направлениях (например, "топот", "око", "заказ").

// Sample Input:

// топот
// Sample Output:

// Палиндром

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var str string

	fmt.Scan(&str)

	str2 := []string{}
	str3 := []string{}

	for _, v := range str {
		str2 = append(str2, string(v))
	}

	for i := utf8.RuneCountInString(str) - 1; i >= 0; i-- {
		str3 = append(str3, string(str2[i]))
	}

	match := "Палиндром"

	for i, v := range str2 {
		if v == str3[i] {
			continue
		} else {
			match = "Нет"
			break
		}
	}

	fmt.Println(match)

}
