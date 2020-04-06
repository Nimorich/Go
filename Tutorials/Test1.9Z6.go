package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num > 0 && num <= 10000 {
		fmt.Print("Positive number. The next number for the number ", num, " is ", num+1, ".\n")
		fmt.Print("Positive number. The previous number for the number ", num, " is ", num-1, ".")
	} else if num < 0 && num >= -10000 {
		fmt.Print("Negative number. The next number for the number ", num, " is ", num+1, ".\n")
		fmt.Print("Negative number. The previous number for the number ", num, " is ", num-1, ".")
	} else if num == 0 {
	}
}
