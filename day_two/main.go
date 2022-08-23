package main

import (
	"fmt"
)

func main() {
	// loops
	for i := 0; i < 10; {
		i = i + 1
		fmt.Println(i)
	}

	fmt.Println()

	// foreach loop
	h := "hello hector"
	for index, word := range h {
		fmt.Printf("%d %c", index, word)
		fmt.Println()
	}
	fmt.Println() // creating an empty line
	// Vectors / Arrays
	var array = [4]int{1, 2, 3, 4}
	// or
	arrayTwo := [...]int{1, 2, 3}
	fmt.Println(array)
	fmt.Println(arrayTwo)
	// empty array
	// emptyArray := []int{}

	// Structs
	type game struct {
		userName string
		weapons  []string
		level    int
	}
	gameOne := game{
		userName: "user1",
		weapons:  []string{"vandale", "phantom"},
		level:    89,
	}
	fmt.Println(gameOne)
	fmt.Println(gameOne.weapons)
	// adding to our weapons vault
	gameOne.weapons = append(gameOne.weapons, "ghost")
	// creating new game struct
	playerPointer := new(game)
	playerPointer.userName = "user2"
	playerPointer.weapons = []string{"ghost"}
	playerPointer.level = 2
	fmt.Println(playerPointer)
	fmt.Println(gameOne.weapons)

	// implementing a function with a struct
	intern1 := intern{name: "nina", level: 1}
	intern1.hey()
	intern1.sayName()
}

type intern struct {
	name  string
	level int
}

// struct function like a impl in Rust
func (intern) hey() {
	fmt.Println("Hello")
}

// we make n a instance of intern so we can use it
func (n intern) sayName() {
	fmt.Println(n.name)
}
