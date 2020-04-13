package main

import "fmt"

func main() {
	array := [5]int{}
	var a, max int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		if a > max {
			max = a
		}
		array[i] = a
	}
	fmt.Print(max)
	// здесь ваш код
	// ...
}
