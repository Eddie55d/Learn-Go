// Напишите программу, которая последовательно делает следующие операции с введённым числом:

// Число умножается на 2;
// Затем к числу прибавляется 100.

// После этого должен быть вывод получившегося числа на экран.

// Sample Input: 1
// Sample Output: 102

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var total int = a*2 + 100
	fmt.Println(total)
}
