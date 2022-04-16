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
}

func sum(sum1 int, sum2 int) int {
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
