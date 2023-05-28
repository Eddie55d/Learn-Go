// Данная задача в основном ориентирована на изучение типа bufio.Reader,
// поскольку этот тип позволяет считывать данные постепенно.

// В тестовом файле, который вы можете скачать из нашего репозитория на github.com,
// содержится длинный ряд чисел, разделенных символом ";". Требуется найти, на какой
// позиции находится число 0 и указать её в качестве ответа. Требуется вывести именно
// позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).

// Например:  12;234;6;0;78 :
// Правильный ответ будет порядковый номер числа: 4

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	s, err := rd.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}

	s = strings.ReplaceAll(s, ";", "\n")

	total := 0

	zero := bufio.NewScanner(strings.NewReader(s))

	for zero.Scan() {
		total++
		if zero.Text() == "0" {
			break
		}
	}

	fmt.Println(total)

}
