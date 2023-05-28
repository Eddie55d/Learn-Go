// Данная задача ориентирована на реальную работу с данными в формате JSON.
// Для решения мы будем использовать справочник ОКВЭД
// (Общероссийский классификатор видов экономической деятельности),
// который был размещен на web-портале data.gov.ru.

// Необходимая вам информация о структуре данных содержится в файле
// structure-20190514T0000.json, а сами данные, которые вам потребуется декодировать,
// содержатся в файле data-20190514T0100.json. Файлы размещены
// в нашем репозитории на github.com.

// Для того, чтобы показать, что вы действительно смогли декодировать документ
// вам необходимо в качестве ответа записать сумму полей global_id всех элементов,
// закодированных в файле.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {

	type myStruct []struct {
		Global int `json:"global_id"`
	}

	file, err := os.Open("./data-20190514T0100.json")
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

	totalNums := 0

	for _, v := range s {
		totalNums += v.Global
	}

	fmt.Println(totalNums)

}
