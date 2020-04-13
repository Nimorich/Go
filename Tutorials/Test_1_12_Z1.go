package main

import "fmt"

func main() {
	var number, tmp, summ int
	fmt.Scan(&number)
	for number > 0 {
		tmp = number % 10
		summ += tmp
		number /= 10
	}
	fmt.Print(summ)
}
