package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	stRun := []rune(str)
	revRun := []rune(str)
	j := 0
	for i := len(stRun) - 1; i >= 0; i-- {
		revRun[j] = stRun[i]
		j++
	}
	if string(stRun) == string(revRun) {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")
	}
}
