package main

import "fmt"

func main() {
	var num, firstNum, finalNum, middleNum int
	fmt.Scan(&num)
	firstNum = num / 100
	middleNum = (num / 10) - (firstNum * 10)
	finalNum = num % 10
	fmt.Println(firstNum, middleNum, finalNum)
	if firstNum != finalNum && finalNum != middleNum && firstNum != middleNum {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
