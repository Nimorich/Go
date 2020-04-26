package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	rn := []rune(text)
	var count int
	if unicode.IsUpper(rn[0]) {
		count++
	}
	if string(rn[len(rn)-1]) == "." {
		count++
	}
	if count == 2 {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
	fmt.Println("Счётчик: ", count)
	fmt.Print("Длинна: ", len(text))
}
