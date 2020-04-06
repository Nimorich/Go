package main

import "fmt"

func main() {
	var x, p, y, year, proc int
	for x <= 0 {
		fmt.Scan(&x)
	}
	for p <= 0 {
		fmt.Scan(&p)
	}
	for y <= 0 {
		fmt.Scan(&y)
	}
	for i := 1; x < y; i++ {
		x *= 100
		proc = (x / 100) * p
		x = x + proc
		x /= 100
		year = i
	}
	fmt.Print(year)
}
