package main

import (
	"fmt"
	"os"
	"strconv"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func PerformOperation(op string, num1, num2 int) int {
	var val int
	switch op {
	case "plus":
		val = Add(num1, num2)
	case "minus":
		val = Subtract(num1, num2)
	default:
		fmt.Println("invalid")
		val = 0
	}

	return val

}

func main() {
	args := os.Args

	operation := args[1]
	num1 := args[2]
	num2 := args[3]

	num1_int, err := strconv.Atoi(num1)
	if err != nil {
		fmt.Println(err)
	}
	num2_int, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(PerformOperation(operation, num1_int, num2_int))
}
