package main

import (
	"fmt"
)

func main() {
	// Hashmaps
	var m map[string]string                 // empty Hashmap
	m = map[string]string{"name": "hector"} // hasmap with values
	fmt.Println(m)
	// another way to add to a hasmap
	m["joey"] = "Not a golang developer"
	fmt.Println(m)
	m["joey"] = "A good Golang developer"
	fmt.Println(m)
	// we can print out just the value
	fmt.Println(m["name"])
	// deleting from our map
	fmt.Println(m)
	delete(m, "joey")
	fmt.Println(m)
	// adding joey back
	m["joey"] = "A good developer"
	// four loop with hasmaps
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Slices
	a := []int{0, 1}
	fmt.Println(a)
	fmt.Println(len(a))

  // Json Conversion
}

type agent {
  name string
  weapon string
  level int
}
