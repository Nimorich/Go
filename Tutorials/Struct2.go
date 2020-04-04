package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var vasya = Person{"Vasya", 45}
	fmt.Println(vasya.age, vasya.name)
	var vasyaPointer *Person = &vasya
	vasyaPointer.age = 29
	fmt.Println(vasya.age)
	(*vasyaPointer).age = 32
	fmt.Println(vasya.age)
}
