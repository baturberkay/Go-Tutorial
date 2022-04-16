package helloWorld

import "fmt"

func HelloWorld() {
	// Uppercase for first letter defines the access level. If function starts with uppercase, this function
	// can be accessed in other modules.

	var name string
	var age int8
	var isStudent bool

	name = "Bob"
	age = 27
	isStudent = true

	fmt.Println("Hello", name)
	fmt.Println("you are", age, "years old.")
	fmt.Println("you are", studentDecider(isStudent))
}

func studentDecider(isStudent bool) string {
	if isStudent {
		return "a student."
	}
	return "not a student."
}
