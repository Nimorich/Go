package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b)
	c = a
	for a > 0 {
		c = a % 10
		e = b
		for e > 0 {
			d = e % 10
			if c == d {
				fmt.Print(c, " ")
			}
			e = e / 10
		}
		a = a / 10
	}
}
