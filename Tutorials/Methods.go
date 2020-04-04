package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

//не изменит оригинальной структуры, для которой вызван
func (p Person) UpdateName(name string) {
	p.Name = name
}

//измеряет оригинальную структуру
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func main() {
	pers := Person{1, "Vasily"}
	pers.SetName("Vasily Romanov")
	//(&pers).SetName("Vasily Romanov")
	fmt.Println("updated person: %#v\n", pers)
	return
	var acc Account = Account{
		Id: 1,
	}
}
