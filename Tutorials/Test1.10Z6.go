package main

import "fmt"

func main() {
	var a int
	for a <= 100 {
		fmt.Scan(&a)
		if a > 100 {
			break
		}
		if a < 10 {
			continue
		}
		fmt.Println(a)
	}
}
