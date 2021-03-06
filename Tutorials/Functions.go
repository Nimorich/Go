package main

import "fmt"

//обычное объявление
func singleIn(in int) int {
	return in
}

//много параметров
func multIn(a, b int, c int) int {
	return a + b + c
}

//именованный результат
func namedReturn() (out int) {
	out = 2
	return
}

//несколько результатов
func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return 0, fmt.Errorf("some error happend")
}

//несколько именованных результатов
func multipleNamedReturn(ok bool) (rez int, err error) {
	if ok {
		err = fmt.Errorf("Some error happend")
		//аналогично return rez, err
		//или return 1, fmt.Errorf("some error happend")
		return
	}
	rez = 2
	return
}

//не фиксированное количество параметров
func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	nums := []int{4, 15, 35, 44, 150, 3, 4, 2}
	fmt.Println(sum(nums...))
	return
}
