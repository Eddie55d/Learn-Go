// Представьте что вы работаете в большой компании где используется модульная архитектура.
// Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу
// какие-то данные. Вы же пишете функцию которая считывает две переменных типа string ,
// а возвращает число типа int64 которое нужно получить сложением этих строк.

// Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того,
// что гоферам платят больше. Поэтому он решил подшутить над вами и подсунул вам свинью.
// Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.

// Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа.
// Под мусором имеются ввиду лишние символы и спец знаки. Разрешается использовать
// только определенные пакеты: fmt, strconv, unicode, strings,  bytes.
// Они уже импортированы, вам ничего импортировать не нужно!

// Считывать и выводить ничего не нужно!

// Ваша функция должна называться adding() !

// Считайте что функция и пакет main уже объявлены!

// Sample Input:

// %^80 hhhhh20&&&&nd

// Sample Output:

// 100

package main

import (
	"fmt"
	"strconv"
)

func preAdding(a string) string {
	var resStr string
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range a {
		var ok bool
		val, err := strconv.Atoi(string(v))
		if err != nil {
			//fmt.Println(err)
			continue
		}
		for _, v := range nums {
			if val == v {
				ok = true
			}
		}
		if ok {
			resStr += string(v)
		}
	}

	return resStr
}

func adding(x, y string) int64 {
	s1 := preAdding(x)
	s2 := preAdding(y)

	resInt1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println(err)
	}

	resInt2, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println(err)
	}

	result := int64(resInt1) + int64(resInt2)

	return result

}

func main() {
	str1 := "%^80"
	str2 := "hhhhh20&&&&nd"

	result := adding(str1, str2)
	fmt.Printf("%d, %T", result, result)
}
