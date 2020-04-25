package main

import "fmt"

func main() {
	var s string = "Это строка"
	for _, b := range s {
		fmt.Print(string(b))
	}
}
