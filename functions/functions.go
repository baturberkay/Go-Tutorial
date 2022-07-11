package functions

import "fmt"

// Sub function without return type
func Sub(x int, y int) int {
	return x - y
}

// Sum function with return type
func Sum(x int, y int) int {
	return x + y
}

// some functions can also return multiple values
func MultipleReturnValues(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

// We can give any number of parameters to the function using variadic functions.
func variadic(params ...int) {
	for _, i := range params {
		fmt.Print(i)
	}
}

func Functions() {
	fmt.Println(Sub(1, 2))
	fmt.Println(Sum(3, 5))
	fmt.Println(MultipleReturnValues(1, 2))
	_, sub := MultipleReturnValues(1, 2) // We can ignore a value using blank identifier (_) symbol.
	fmt.Println(sub)
	variadic(1, 2, 3, 4, 5, 6, 7, 8)
}
