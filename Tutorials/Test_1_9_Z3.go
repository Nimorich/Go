package main

import "fmt"

func main() {
	var a, result int
	fmt.Scan(&a)
	if 10000 > a && a > 1000 {
		result = a / 1000
		fmt.Println(result)
	} else if 1000 > a && a > 100 {
		result = a / 100
		fmt.Println(result)
	} else if 100 > a && a > 10 {
		result = a / 10
		fmt.Println(result)
	} else if 10 > a && a > 0 {
		result = a
		fmt.Println(result)
	} else if 10000 == a {
		result = a / 10000
		fmt.Println(result)
	}
}
