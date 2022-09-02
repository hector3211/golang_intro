package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func newPerson(name string, age int) *Person {
	p := Person{Name: name, Age: age}
	return &p
}

func main() {
	fmt.Println("hello hector")
	fmt.Println(newPerson("hector", 26))
}
