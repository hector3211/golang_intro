package main

import (
	"encoding/json"
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
	sFrom := `{"name":"madDog","weapon":{"weapon_name":"headhunter","weapon_level":99},"level":1}`
	fmt.Println(sFrom)
	var chamber Agent // referring to our Agent struct
	err := json.Unmarshal([]byte(sFrom), &chamber)
	if err != nil {
		fmt.Println("bad stuff")
	}
	fmt.Println(chamber)
	// Go Object to Json
	sTo, err := json.Marshal(chamber)
	if err != nil {
		fmt.Println("something bad")
		fmt.Println(err)
	}
	fmt.Println(sTo) // right now this returns bytes
	// converting it to a string
	fmt.Printf("%s\n", sTo) // %s tells we want it as Json
}

type Agent struct { // CAP first letter
	Name   string `json:"name"`   // CAP first letter
	Weapon Weapon `json:"weapon"` // CAP first letter
	Level  int    `json:"level"`  // CAP first letter
}
type Weapon struct {
	Name  string `json:"weapon_name"`
	Level int    `json:"weapon_level"`
}
