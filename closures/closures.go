package closures

import "fmt"

// closures are anonymous functions of Go.
// this function will return the result of the anonymous function.
func closureExample() func() int {
	closureCounter := 0
	//It can be defined without a name.
	return func() int {
		closureCounter += 1
		return closureCounter
	}
}

func Closure() {
	// closureCounter keeps the last value of the function and updates its value when it called.
	closureCounter := closureExample()
	fmt.Println(closureCounter()) // prints 1
	fmt.Println(closureCounter()) // prints 2
	fmt.Println(closureCounter()) // prints 3

	// new function has a new value.
	newClosureCounter := closureExample()
	fmt.Println(newClosureCounter()) // prints 1

}
