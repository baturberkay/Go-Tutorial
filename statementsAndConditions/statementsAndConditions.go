package statementsAndConditions

import "fmt"

func StatementsAndConditions() {
	// Go has no ternary operator because of readability of the code. Classical if/else statements need to be used.

	const evenNumber int = 8
	const oddNumber int = 9

	if evenNumber%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if oddNumber < 6 {
		fmt.Println("less than 6")
	} else if oddNumber > 6 {
		fmt.Println("greater than 6")
	} else {
		fmt.Println("equals to 6")
	}

}
