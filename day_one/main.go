// package main lets go know we'll be executing this
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt package allows us to use Print
	fmt.Println("hello world!")

	// Creating varibles
	var firstVariable string = "Our first variable"
	var number int = 10
	// another way of creating a varible
	secondVariable := "Second variable"
	fmt.Println(firstVariable, number, secondVariable)

	// types
	integer := 1
	float := 1.0
	aBool := true
	aString := "string"
	fmt.Println(integer, float, aBool, aString)

	// Handling strings
	s := "Hello World!"
	// indexing strings
	fmt.Println(s[:1]) // returns letter H
	// Concatening strings
	s += " Again"
	fmt.Println(s)
	// upper case strings
	newString := "hello"
	fmt.Println(strings.ToUpper(newString))
	fmt.Println(strings.ToLower(newString))
	// conversion between strings
	// converting a number to a string
	x := 123
	y := strconv.Itoa(x)
	fmt.Printf("%T", y) // retruns its a String

	// declaring a constant
	const numberFive int = 5
	// multiple constants
	const (
		a = 1
		b = 2
	)
}
