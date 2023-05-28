// Напишите функцию которая принимает канал и число N типа int.
// Необходимо вернуть значение N+1 в канал.
// Функция должна называться task().

package main

import (
	"fmt"
)

func task(N int, ch chan int) {
	ch <- N + 1
}

func main() {
	ch := make(chan int)
	a := 2
	go task(a, ch)
	fmt.Println(<-ch)
}
