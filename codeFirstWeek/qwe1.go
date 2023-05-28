package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 3
	}()
	fmt.Print(<-ch)

}
