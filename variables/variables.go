package variables

import "fmt"

const constantVariable string = "This is a constant variable." // const keyword is used to declare constant variables.

func Variables() {

	var one, two = 1, 2 // one or more variables can be declared using var keyword.
	//one,two := 1,2 is another way to declare one or more variables
	//var one, two int = 1,2 type (int) can be omitted
	fmt.Println(one, two)

	var greeting = "Hi Bob!"
	fmt.Println(greeting)

	var varWithoutInitialValue int // go will set initial value as 0.
	fmt.Println("varWithoutInitialValue:", varWithoutInitialValue)

	varWithInitialValue := 2 // := operator is shorthand for declaring and initializing a variable.
	fmt.Println("varWithInitialValue:", varWithInitialValue)

	fmt.Println("constantVariable: ", constantVariable)

}
