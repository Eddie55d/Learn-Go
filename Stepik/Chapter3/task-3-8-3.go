// Напишите элемент конвейера (функцию), что запоминает предыдущее значение и
// отправляет значения на следующий этап конвейера только если оно отличается от того,
// что пришло ранее.

// Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы
// будете получать строки, во второй вы должны отправлять значения без повторов.
// В итоге в outputStream должны остаться значения, которые не повторяются подряд.
// Не забудьте закрыть канал.

// Функция должна называться removeDuplicates()

// Выводить или вводить ничего не нужно!

package main

import (
	"fmt"
)

func removeDuplicates(inputStream, outputStream chan string) {
	c := make([]string, 1)

	for v := range inputStream {
		if v != c[0] {
			outputStream <- v
			c[0] = v
		}
	}
	defer close(outputStream)
}

func main() {
	inputStream := make(chan string, 15)
	outputStream := make(chan string)

	str := "qwee1233"
	fmt.Println(str)

	go func() {
		for _, v := range str {
			inputStream <- string(v)
		}
		defer close(inputStream)
	}()

	go removeDuplicates(inputStream, outputStream)
	for i := range outputStream {
		fmt.Printf(i)
	}

}
