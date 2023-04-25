// Дано натуральное число, выведите его последнюю цифру.

// Формат входных данных
// На вход дается натуральное число N, не превосходящее 10000.

// Формат выходных данных
// Выведите одно целое число - ответ на задачу.

// Sample Input: 123
// Sample Output:   3

package main

import (
	"fmt"
	"strconv"
)

var (
	num int
	str string
)

func main() {
	fmt.Print("Введите число не более 10000: ")
	fmt.Scan(&num)
	if num <= 10000 {

		str := strconv.Itoa(num)
		fmt.Println(str)
		s := str[len(str)-1]
		//fmt.Println(string(s))
		end, _ := strconv.Atoi(string(s))
		fmt.Printf("%v the last digit in the number %v", end, num)
	} else {
		println("число больше 10000")
	}
}
