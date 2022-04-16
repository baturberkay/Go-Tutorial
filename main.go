package main

import "fmt"

func main() {

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
