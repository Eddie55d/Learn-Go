// На стандартный ввод подается строковое представление двух дат, разделенных запятой
// (формат данных смотрите в примере).

// Необходимо преобразовать полученные данные в тип Time, а затем вывести
// продолжительность периода между меньшей и большей датами.

// Sample Input:

// 13.03.2018 14:00:15,12.03.2018 14:00:15

// Sample Output:

// 24h0m0s

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	str := "13.03.2018 14:00:15,12.03.2018 14:00:15"

	subStr := strings.Split(str, ",")

	time1, err := time.Parse("02.01.2006 15:04:05", subStr[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	time2, err := time.Parse("02.01.2006 15:04:05", subStr[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case time1.After(time2):
		fmt.Println(time1.Sub(time2))
	default:
		fmt.Println(time2.Sub(time1))
	}
}
