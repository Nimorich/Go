package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for a > 0 {
		fmt.Print(a%10)
		a /= 10
	}
}
