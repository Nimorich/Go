package main

import "fmt"

func main() {
	fmt.Print(fibonacci(3))
}

func fibonacci(n int) int {
	a := 1
	b := 1
	c := 0
	for i := 0; i < n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}
