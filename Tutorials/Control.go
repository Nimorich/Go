package main

import "fmt"

func main() {
	//простое условие
	var boolVal bool = true
	if boolVal {
		fmt.Println("boolVal is true")
	}

	mapVal := map[string]string{"name": "rvasily", "name2": "kruosani", "name3": "perosani"}
	//условие с блоком инициализации
	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name =", keyValue)
	}

	//получаем только признак существования ключа
	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println("key 'name' exist")
	}

	cond := 4
	//множественные if else
	if cond == 1 {
		fmt.Println("cond is 1")
	} else if cond == 2 {
		fmt.Println("cond is 2")
	} else {
		fmt.Println("другое")
	}

	//switch по 1 переменной
	strVal := "name1"
	switch strVal {
	case "name":
		fmt.Println("Проваливание")
		fallthrough
	case "test", "lastName":
		fmt.Println("next case")
		//some work
	default:
		fmt.Println("default case")
		fmt.Println("default case2")
		//some work
	}

	//switch как замена многим if else
	var val1, val2 = 2, 2
	switch {
	case val1 > 1 || val2 < 11:
		fmt.Println("first block")
	case val2 > 10:
		fmt.Println("second block")
	}

	//выход из цикла
	//Loop:
	for key, val := range mapVal {
		println("switch in loop", key, val)
		switch {
		case key == "lastName":
			println("dont pront this")
			break
		case key == "firstName" && val == "Vasily":
			println("switch - break loop here")
			break //Loop
		} //конец for
	}

}
