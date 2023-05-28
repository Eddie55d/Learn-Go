// Дается строка. Нужно удалить все символы, которые встречаются более
// одного раза и вывести получившуюся строку

// Sample Input:   zaabcbd
// Sample Output:  zcd

package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string
	fmt.Scan(&str)

	str2 := []string{}
	str3 := []string{}
	str4 := []string{}

	for _, v := range str {
		s := string(v)
		str2 = append(str2, s)
		str3 = append(str3, s)
	}

	for _, v := range str2 {
		count := 0
		for i := 0; i < len(str2); i++ {
			if v == str3[i] {
				count++
			}
		}
		if count == 1 {
			str4 = append(str4, v)
		}
	}

	fmt.Println(strings.Join(str4, ""))

}
