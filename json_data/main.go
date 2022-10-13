package main

import (
	"encoding/json"
	"fmt"
)

type Pokemon struct {
	Name     string `json:name`
	Level    int    `json:level`
	Type     string `json:type`
	Weakness string `json:weakness`
}

func main() {
	jsonData := (`
    {
        "name":"charmander",
        "level":34,
        "type":"fire",
        "weakness":"water",
        "attacks":["fireball","DragonsBreath"]
    }
    `)

	var jsonPoke Pokemon

	// checking json data if it is json data
	checkValid := json.Valid([]byte(jsonData))

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal([]byte(jsonData), &jsonPoke)
		fmt.Printf("%#v\n", jsonPoke)
	} else {
		fmt.Println("Json was not valid")
	}

	test()
}

func test() {

	jsonData := []byte(`
    {
        "name":"charmander",
        "level":34,
        "type":"fire",
        "weakness":"water",
        "attacks":["fireball","DragonsBreath"]
    }
    `)
	var newJosnData map[string]interface{}
	json.Unmarshal(jsonData, &newJosnData)
	fmt.Printf("%#v\n", newJosnData)

	for k, v := range newJosnData {
		fmt.Printf("key is: %v and values is: %v and Type is: %T\n", k, v, v)
	}
}
