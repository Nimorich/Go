package main

import "fmt"

func main() {
	var k int
	var h, m int
	fmt.Scan(&k)
	h = (k / 3600)
	m = (k / 60) - (h * 60)
	fmt.Print("It is ", h, " hours ", m, " minutes.")
}
