package arrays

import "fmt"

func Arrays() {

	var arr [10]int      // declaring array without user defined  initial values. Go fills the array with 0s. [0,0,0,0,..,0]
	arr1 := [2]int{1, 2} // declaring array with user defined initial values

	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	fmt.Println(arr1)

}
