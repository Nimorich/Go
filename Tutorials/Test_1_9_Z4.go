package main

import "fmt"

func main() {
	var num string
	var a, b, c, d, e, f, sum1, sum2 int
	fmt.Scan(&num) //шестизначное число
	a = int(num[0])
	b = int(num[1])
	c = int(num[2])
	d = int(num[3])
	e = int(num[4])
	f = int(num[5])
	sum1 = a + b + c
	sum2 = d + e + f
	if sum1 == sum2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
