// На вход дается строка, из нее нужно сделать другую строку,
// оставив только нечетные символы (считая с нуля)

// Sample Input:   ihgewlqlkot
// Sample Output:  hello

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Scan(&str)

	str2 := []string{}

	for i, v := range str {
		if i%2 != 0 {
			str2 = append(str2, string(v))
		}
	}

	fmt.Println(strings.Join(str2, ""))

}
