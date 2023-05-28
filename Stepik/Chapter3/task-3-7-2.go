// На стандартный ввод подается строковое представление даты и времени определенного
// события в следующем формате:

// 2020-05-15 08:00:00

// Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести
// дату на стандартный вывод в том же формате.

// Если же событие должно произойти после обеда, необходимо перенести его на то же
// время на следующий день, а затем вывести на стандартный вывод в том же формате.

// Sample Input:

// 2020-05-15 08:00:00

// Sample Output:

// 2020-05-15 08:00:00

package main

import (
	"fmt"
	"time"
)

func main() {

	s := "2020-05-15 15:00:00"

	dinner := "13:00:00"

	dinnerTime, _ := time.Parse(time.TimeOnly, dinner)

	timeEvent, _ := time.Parse(time.DateTime, s)

	onlyTimeEvent, _ := time.Parse(time.TimeOnly, timeEvent.Format(time.TimeOnly))

	switch {
	case onlyTimeEvent.After(dinnerTime):
		eventForTomorrow := timeEvent.AddDate(0, 0, 1)
		fmt.Println(eventForTomorrow.Format(time.DateTime))
	default:
		fmt.Println(timeEvent.Format(time.DateTime))
	}

}
