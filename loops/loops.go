package loops

import "fmt"

func Loops() {

	i := 0
	for i <= 10 { // for with single condition. Like while loop
		fmt.Println(i)
		i++ // i = i + 1
	}

	for j := 10; j <= 15; j++ {
		fmt.Println(j)
	}

	for true { // while true
		if i < 15 {
			fmt.Println("always true")
			i++
			continue //continue to the next iteration using continue keyword
		}
		break
	}
}
