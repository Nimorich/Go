// Последовательность состоит из натуральных чисел
//и завершается числом 0.Определите количество элементов
//этой последовательности, которые равны ее наибольшему элементу.
//Формат входных данных
//Вводится непустая последовательность натуральных чисел,
//оканчивающаяся числом 0 (само число 0 в последовательность не входит,
//а служит как признак ее окончания).

package main

import "fmt"

func main() {
	var a int = 1
	var b, count int
	for a != 0 {
		fmt.Scan(&a)
		if a != 0 {
			//fmt.Println(a)
			if a > b {
				b = a
				count = 0
			}
			if a == b {
				count++
			}
		}
	}
	fmt.Print(count)
}
