package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Height int
	Weight int
}

func main() {
	// check if program is working
	fmt.Println("Hellow hector")
	// create a new person from our Person struct
	newPerson := &Person{
		Name:   "hector",
		Age:    28,
		Height: 6,
		Weight: 170,
	}
	fmt.Println(newPerson)
	// call our calculate function that calculates two numbers
	result, err := calculate("+", 1, 1)
	// retrun an error if there is
	if err != nil {
		fmt.Println("oops something happened")
	} else {
		fmt.Printf("result is :%v", result) // return result
	}
}

// function that calculates two number based on an operation
// and returns an error if somrthing is wrong
func calculate(op string, numberOne int, numberTwo int) (int, error) {
	var equal int
	if op == "-" {
		equal = numberOne - numberTwo
	} else if op == "+" {
		equal = numberOne + numberTwo
	} else if op == "*" {
		equal = numberOne * numberTwo
	} else if op == "/" {
		equal = numberOne / numberTwo
	} else {
		return equal, fmt.Errorf("Please provide a valid operation")
	}
	return equal, nil
}
