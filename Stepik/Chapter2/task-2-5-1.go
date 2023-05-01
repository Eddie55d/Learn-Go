// На вход подается строка. Нужно определить, является ли она правильной или нет.
// Правильная строка начинается с заглавной буквы и заканчивается точкой.
// Если строка правильная - вывести Right иначе - вывести Wrong

// Маленькая подсказка: fmt.Scan() считывает строку до первого пробела,
// вы можете считать полностью строку за раз с помощью bufio:

// text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

// Sample Input:  Быть или не быть.

// Sample Output:  Right

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	rs := []rune{}

	for _, v := range text {
		rs = append(rs, v)
	}

	last := rs[(len(rs) - 3)]

	if string(last) == "." && unicode.IsUpper(rs[0]) {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

}
