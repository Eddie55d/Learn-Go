// Определите является ли билет счастливым. Счастливым считается билет,
// в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

// Формат входных данных

// На вход подается номер билета - одно шестизначное  число.

// Формат выходных данных
// Выведите "YES", если билет счастливый, в противном случае - "NO".

// Sample Input:  613244
// Sample Output: YES

package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Scan(&a)

	n6 := int(math.Mod(a, 10))
	n5 := int(math.Mod((a / 10), 10))
	n4 := int(math.Mod((a / 100), 10))
	n3 := int(math.Mod((a / 1000), 10))
	n2 := int(math.Mod((a / 10000), 10))
	n1 := int(math.Mod((a / 100000), 10))
	//fmt.Println(n6, n5, n4, n3, n2, n1)
	leftNum := n1 + n2 + n3
	//fmt.Println(leftNum)
	rightNum := n4 + n5 + n6
	//fmt.Println(rightNum)
	if leftNum == rightNum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
