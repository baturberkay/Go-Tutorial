package arithmeticOperations

import "fmt"

func ArithmeticOperations() {

	sumResult := sum(1, 2) // declaring variable without specifying variable type
	fmt.Println("sum:=1+2=>", sumResult)

	divisionResult := division(4, 2)
	fmt.Println("division:=2/4=>", divisionResult)

	subtractionResult := subtraction(8, 2)
	fmt.Println("subscription:= 8-2=>", subtractionResult)

	productionResult := production(2, 6)
	fmt.Println("production:= 2*6=>", productionResult)

	// functions can multiple two values
	res1, res2 := multipleReturn(2, 6)
	fmt.Println("res1, res2", res1, res2)
}

// more than one parameter of the same type can be written in a single variable type.
func sum(sum1, sum2 int) int {
	return sum1 + sum2
}

func division(dividend int, divisor int) float32 {
	if divisor == 0 {
		return 0
	}
	return float32(dividend / divisor)
}

func subtraction(subs1 int, subs2 int) int {
	return subs1 - subs2
}

func production(prod1 int, prod2 int) int {
	return prod1 * prod2
}

func multipleReturn(num1 int, num2 int) (int, int) {
	sum1 := num2 + num1
	prod1 := num1 * num2

	return sum1, prod1
}
