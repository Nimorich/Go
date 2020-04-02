package main

import "fmt"

func main() {
	//динамический массив Слайс
	/*var buf0 []int             //len=0, cap=0
	buf1 := []int{}            //len=0, cap=0
	buf2 := []int{42}          //len=1, cap=1
	buf3 := make([]int, 0)     //len=0, cap=0
	buf4 := make([]int, 5)     //len=5, cap=5
	buf5 := make([]int, 5, 10) //len=5, cap=10

	//обращение к элементам
	someInt := buf2[0]*/
	//ошибка при выполнении
	//panic: runtime error: index out of range
	//someOtherInt := buf2[1]

	//добавление элементов
	var buf []int            //len=0, cap=0
	buf = append(buf, 9, 10) //len=2, cap=2
	buf = append(buf, 12)    //len=3 cap=4 потому что множится(x2)
	//добавление другого слайса
	otherBuf := make([]int, 3)     //[0,0,0]
	buf = append(buf, otherBuf...) //len=6, cap=8
	var bufLen, bufCap int = len(buf), cap(buf)
	fmt.Println(bufLen, bufCap)
}
