package functions

import "fmt"

// function without return type
func add(x int, y int) {
	fmt.Println(x + y)
}

// function with return type
func sum(x int, y int) int {
	return x + y
}

//functions can also return multiple values
func multipleReturnValues() (int, int) {
	a := 2
	b := 3
	sum := a + b
	sub := a - b
	return sum, sub
}

func Functions() {
	add(1, 2)
	fmt.Println(sum(3, 5))
	fmt.Println(multipleReturnValues())
	_, sub := multipleReturnValues() // We can ignore a value using blank identifier (_) symbol.
	fmt.Println(sub)
}
