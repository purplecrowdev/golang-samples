package main

import (
	"fmt"
	"strings"
)

func main() {

	// check if the string has prefix

	const COUNTRY = "INDIA"

	doesStartWithI := strings.HasPrefix(COUNTRY, "I")
	fmt.Println(doesStartWithI)

	// replace string

	nameDesc := "my name is akshat"
	replacedString := strings.Replace(nameDesc, "is", "am", -1)
	fmt.Println(replacedString)

	// Lower or Upper

	fmt.Println(strings.ToLower(COUNTRY))

	// repeat a string

	word := "a"
	fmt.Println(strings.Repeat(word, 10))

}
