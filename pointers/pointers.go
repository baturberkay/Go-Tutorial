package pointers

import "fmt"

func Pointers() {
	value := 123
	fmt.Println("initial value: ", value)

	resetVal(value)
	fmt.Println("reset value: ", value)

	resetPtr(&value)
	fmt.Println("reset ptr:", value)

	fmt.Println("value of the pointer: ", &value)
}

func resetVal(value int) {
	value = 0
}

func resetPtr(ptr *int) {
	*ptr = 0
}
