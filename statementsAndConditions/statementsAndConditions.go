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

	switch { // switch without an expression can be used like if/else statements
	case evenNumber < 6:
		fmt.Println("less than 6")
	case evenNumber > 6:
		fmt.Println("greater than 6")
	default:
		fmt.Println("equals to 6")
	}

	number := 5
	switch number {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("not recognized")
	}

}
