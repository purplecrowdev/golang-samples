package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num int

	fmt.Println(num) // 0

	name := "akshat"
	fmt.Println(reflect.TypeOf(name)) // string

	const (
		one = 1
		two = 2
	)

	fmt.Println(one, two) //1, 2

	f, g := 10, 12 // multiple assigns

	fmt.Println(f, g)
}
