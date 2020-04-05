package main

import "fmt"

func main() {
	var ear int
	fmt.Scan(&ear)
	if ear%4 == 0 && ear%100 != 0 || ear%400 == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
