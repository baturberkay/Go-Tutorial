package slices

import "fmt"

func Slices() {

	slice := make([]int, 3) // Go initially fills the slice with 0s.
	fmt.Println("slice:", slice)
	fmt.Println("length of slice:", len(slice))

	// setting a new value to an element
	slice[0] = 0
	slice[1] = 1
	slice[2] = 2

	// getting a value of an element
	fmt.Println("slice[0]", slice[0])
	fmt.Println("slice[1]", slice[1])
	fmt.Println("slice[2]", slice[2])

	//appending one or more elements
	slice = append(slice, 4, 5)

	fmt.Println(slice)

}
