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

	//  keep the data type in ming when assigning the const

	const (
		Pi               = 3.1456777
		Avogardo float32 = 6.022e23
	)

	fmt.Println(Pi)
	fmt.Println(reflect.TypeOf(Pi))
	fmt.Println(Avogardo)
	fmt.Println(reflect.TypeOf(Avogardo))

	f, g := 10, 12 // multiple assigns

	fmt.Println(f, g)
}
