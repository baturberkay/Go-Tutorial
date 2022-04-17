package maps

import "fmt"

func Maps() {
	maps := make(map[int]string)               // creating a map using make keyword.
	newMap := map[string]string{"name": "bob"} // creating map with initial values.

	maps[1] = "one" // adding key/value pair to the map
	maps[2] = "two"
	maps[3] = "three"

	for i := 1; i <= len(maps); i++ {
		fmt.Println(maps[i]) // get value using key
	}

	delete(maps, 2) // delete a pair from map using key
	fmt.Println(newMap["name"])

	// found returns true if the key exists.
	// otherwise, v will be 0 and exist will be false
	if v, exist := maps[1]; exist {
		fmt.Println(v, exist) // one true
	}

	for key, value := range maps { // iterate over maps using range
		fmt.Println(key, value)
	}
}
