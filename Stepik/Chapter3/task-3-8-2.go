// Напишите функцию которая принимает канал и строку.
// Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.

// Функция должна называться task2().

package main

import (
	"fmt"
)

func task2(s string, ch chan string) {
	for i := 1; i <= 5; i++ {
		ch <- s + " "
	}
}

func main() {
	ch := make(chan string)
	a := "hello"
	go task2(a, ch)
	fmt.Printf("%q", <-ch)
}
