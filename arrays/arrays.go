package arrays

import "fmt"

func Arrays() {

	var arr [10]int      // declaring array without initial values
	arr1 := [2]int{1, 2} // declaring array with initial values

	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	fmt.Println(arr1)

}
