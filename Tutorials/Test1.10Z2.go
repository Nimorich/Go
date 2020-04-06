//Требуется написать программу, при выполнении которой с клавиатуры считываются
// два натуральных числа A и B (каждое не более 100, A < B).
// Вывести сумму всех чисел от A до B  включительно.
package main

import (
	"fmt"
)

func main() {
	var a, b, sum int
	fmt.Scan(&a, &b)
	for a > 100 || a <= 0 {
		fmt.Scan(&a)
	}
	for b > 100 || b <= 0 {
		fmt.Scan(&b)
	}
	if a > b {
		temp := a
		a = b
		b = temp
	}
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Print(sum)
}
