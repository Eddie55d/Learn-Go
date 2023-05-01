//  Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные
//  требования. Длина пароля должна быть не менее 5 символов, он должен содержать
//  только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль.
//  Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

// Sample Input:   fdsghdfgjsdDD1
// Sample Output:  Ok

package main

import (
	"fmt"
	"strings"
)

func main() {
	var password string
	fmt.Scan(&password)

	charSet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	str2 := []string{}

	for _, v := range password {
		s := string(v)
		str2 = append(str2, s)
	}

	correct := true

	for _, v := range password {

		if strings.Contains(string(charSet), string(v)) == true {
			continue
		} else {
			correct = false
		}
	}

	if correct && len(str2) >= 5 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}

}
