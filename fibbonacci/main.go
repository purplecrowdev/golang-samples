package main

import "fmt"

func fib(n int) int {
	mem := make([]int, n)

	mem[0] = 1
	mem[1] = 1

	for i := 2; i < n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}

	return mem[n-1]
}

func main() {

	fmt.Println(fib(5))

}
