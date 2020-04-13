package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n)
	var array []int
	for i := 0; i <= n-1; i++ {
		fmt.Scan(&s)
		array = append(array, s)
	}
	for i, value := range array {
		if i%2 == 0 {
			fmt.Print(value, " ")
		}
	}
}
