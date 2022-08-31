package main

import (
	"errors"
	"fmt"
)

func main() {
	sayHello()
	fmt.Println(sum(1, 1))
	fmt.Println(nameLenght("hector"))
	value, error := result("")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(value)
	}
	num := 1
	fmt.Println(num)
	withOut(num)
	fmt.Println(num)
	withPointer(&num)
	fmt.Println(num)
}

// functions
func sayHello() {
	fmt.Println("hello!")
}

// function that returns a integer
func sum(a, b int) int {
	return a + b
}

// function that returns two values
func nameLenght(name string) (string, int) {
	return name, len(name)
}

// function that returns a value or an error
func result(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("emppty name param")
	} else {
		return "hello " + name, nil
	}
}

// function with pointer
func withPointer(num *int) {
	*num += 1
}

// function without pointer
func withOut(num int) {
	num += 1
}
