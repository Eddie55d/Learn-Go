// Используем анонимные функции на практике.

// Внутри функции main (объявлять ее не нужно) вы должны объявить
// функцию вида func(uint) uint, которую внутри функции main в дальнейшем можно
// будет вызвать по имени fn.

// Функция на вход получает целое положительное число (uint), возвращает
// число того же типа в котором отсутствуют нечетные цифры и цифра 0.
// Если же получившееся число равно 0, то возвращается число 100.
// Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

// Пакет main объявлять не нужно!
// Вводить и выводить что-либо не нужно!
// Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

// Sample Input:    727178

// Sample Output:   28

package main

import (
	"fmt"
	"strconv"
)

func anon() func(uint) uint {

	return func(a uint) uint {
		var numsStr string

		b := strconv.Itoa(int(a))

		for _, v := range b {
			d, _ := strconv.Atoi(string(v))
			if d%2 == 0 && d != 0 {
				numsStr += string(v)
			}
		}

		numsUint, _ := strconv.ParseUint(numsStr, 10, 64)

		c := uint(numsUint)

		if c == 0 {
			c = 100
		}
		return c
	}

}

func main() {
	var num uint
	fmt.Scan(&num)

	fn := anon()
	fmt.Println(fn(num))

}
