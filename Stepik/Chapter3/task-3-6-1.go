// На стандартный ввод подаются данные о студентах университетской группы
// в формате JSON:

// {
//     "ID":134,
//     "Number":"ИЛМ-1274",
//     "Year":2,
//     "Students":[
//         {
//             "LastName":"Вещий",
//             "FirstName":"Лифон",
//             "MiddleName":"Вениаминович",
//             "Birthday":"4апреля1970года",
//             "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//             "Phone":"+7(948)709-47-24",
//             "Rating":[1,2,3]
//         },
//         {
//             // ...
//         }
//     ]
// }

// В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
// Требуется прочитать данные, и рассчитать среднее количество оценок,
// полученное студентами группы. Ответ на задачу требуется записать на стандартный вывод
// в формате JSON в следующей форме:

// {
//     "Average": 14.1
// }

// Как вы понимаете, для декодирования используется функция Unmarshal, а для кодирования
// MarshalIndent (префикс - пустая строка, отступ - 4 пробела).

// Если у вас возникли проблемы с чтением / записью данных, то этот комментарий для вас:
// в уроках об интерфейсах и работе с файлами мы рассказывали,
// что стандартный ввод / вывод - это файлы, к которым можно обратиться
// через os.Stdin и os.Stdout соответственно, они удовлетворяют
// интерфейсам io.Reader и io.Writer, из них можно читать, в них можно писать.

// Один из способов чтения, использовать ioutil.ReadAll:

// data, err := ioutil.ReadAll(os.Stdin)
// if err != nil {
//     // ...
// }

// // data - тип []byte
// Задачу можно выполнить и другими способами, в частности использовать bufio.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {

	type myStruct struct {
		Students []struct {
			Rating []int `json:"Rating"`
		} `json:"Students"`
	}

	file, err := os.Open("./students.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	rd, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	var s myStruct

	if err := json.Unmarshal(rd, &s); err != nil {
		fmt.Println(err)
		return
	}

	numStudents := len(s.Students)
	totalNums := 0

	for _, v := range s.Students {
		for range v.Rating {
			totalNums++
		}
	}

	type averageAmount struct {
		Average float32
	}

	total := averageAmount{float32(totalNums) / float32(numStudents)}

	data, err := json.MarshalIndent(total, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = os.Stdout.Write(data)

}
