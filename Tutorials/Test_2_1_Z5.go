package main

import "fmt"

func main() {
	a, b := sumInt(1, 0, 6, 8, 2)
	fmt.Println(a, b)
}
func sumInt(ar ...int) (int, int) {
	var sum, count int
	for _, value := range ar {
		sum += value
		count++
	}
	return count, sum
}
