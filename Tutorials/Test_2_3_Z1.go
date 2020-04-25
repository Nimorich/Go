package main

import (
	"fmt"
)

func zero(x *int) {
	*x = 2 * *x
}
func main() {
	x := 11
	zero(&x)
	fmt.Print(x)
}
