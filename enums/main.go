package main

import "fmt"

type Season int64

const (
	Summer Season = iota
	Rainy
	Winter
	Autumn
)

func (s Season) String() string {
	switch s {
	case Summer:
		return "summer"
	case Rainy:
		return "rainy"
	case Winter:
		return "winter"
	case Autumn:
		return "autumn"
	}
	return "unknown"
}

func main() {

	// type Season int

	// const (
	// 	Summer Season = iota
	// 	Rainy
	// 	Winter
	// 	Autumn
	// )

	// fmt.Println(Winter)

	fmt.Println(Rainy)

}
