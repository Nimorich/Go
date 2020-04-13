package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var x []int
	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		x = append(x, num)
	}
	fmt.Print(x[3])
}
