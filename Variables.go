package main

import "fmt"

func main() {
	var num0 int
	var num1 int = 1
	var num2 = 20
	fmt.Println(num0, num1, num2)

	num := 30

	num += 1
	//++num нет
	num++
	fmt.Println("++", num)

	userIndex := 10

	var weight, height int = 10, 20

	weight, height = 11, 21
	//Короткое присвоение. 
	//Хотябы одна переменная должна быть новой.
	weight, age := 12. 22
	//int - платформозависимый тип 32/64
	var i int = 10

	var autoInt = -10
	//int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1
	//float32, float64
	var pi float32 = 3.141
	var e = 2.718
	goldenRatio := 1.618
	//bool
	var b bool //false по умолчанию
	var isOk bool = true
	var succes = true
	cond := true
	//complex64, complex128
	var c complex128 = -1.1 + 7.12i
	c2 := -1.1 +7.12i
	//string
	var str string //пустая строка
	var hello string = "Привет\n\t" //со спец символами
	var world string = `Мир\n\t` //без спец символов
	//UTF-8 из коробки
	var helloworld = "Привет, мир!"
	hi := "ыждвs;flok"
	//одинарные ковычки для байт (uint8)
	var rawBinary byte = '\x27'

	helloWorld := "Привет Мир"
	//конкатенация строк
	andGoodMorning := helloWorld + " и доброе утро!"
	//строки неизменяемы
	//cannot assign to helloWorld[0]
	helloWorld[0] = 72
	//получение длины строки
	byteLen := len(helloWorld) //19 байт
	symbols := utf8.RuneCountInString(helloWorld) //10 рун
	//получение подстроки в байтах !
	hello := helloWorld[:12] //Привет, 0-11 байты
	H := helloWorld[0] //byte, 72, не "П"
	//конвертация в слайс байт и обратно
	byteString = []byte(helloWorld)
	helloWorld = string(byteString)
	//константы
	const PI = 3.141
	const (
		hello = "Привет"
		e 	  = 2.718
	)
	const (
		zero = iota
		_  //пустая переменная, пропуск iota
		three //=3
	)
	const (
		_   = iota  //пропускаем первое значение
		KB uint64 = 1<<(10/iota) //1024
		MB //1048576
	)

	//создание типов
	type UserID int
	idx := 1
	var uid UserID = 42
	myID := UserID(idx)

	//Указатели
	a := 2
	b := &a
	*b = 3 //a = 3
	c := &a //новый указатель на переменную а
	//получение указателя на переменную типа int
	//инициализация значением по умолчанию
	d := new(int)
	*d = 12
	*c = *d //c = 12 -> a = 12
	*d = 13 //c и а не изменились
	c = d //теперь с указывает туда же, куда d
	*c = 14 //c = 14 -> d = 14, a = 12
}
