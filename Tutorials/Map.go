package main

import "fmt"

func main() {

	var userMap map[string]string = map[string]string{"key": "value", "key2": "value"}
	fmt.Println(userMap["key2"])
}
