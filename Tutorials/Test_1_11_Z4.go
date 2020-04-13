package main

import "fmt"

func main() {
	var n, s, count int
	fmt.Scan(&n)
	var array []int
	for i := 0; i <= n-1; i++ {
		fmt.Scan(&s)
		array = append(array, s)
	}
	for _, value := range array {
		if value > 0 {
			count++
		}
	}
	fmt.Print(count)
}
