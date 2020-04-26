package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)
	r := []rune(x)
	var str []rune
	for i, v := range r {
		if i%2 != 0 {
			str = append(str, v)
		}
	}
	fmt.Print(string(str))
}
