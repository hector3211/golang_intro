package controllers

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const Url = "https://pokeapi.co/api/v2/pokemon/"

func GetValue(c *fiber.Ctx) error {
	return c.SendString("value is:" + c.Params("value"))
}

type Pokemons struct {
	Name string `json:name`
	Url  string `json:url`
}
type Results struct {
	Results []Pokemons `json:results`
}

func GetAllPokemons(c *fiber.Ctx) error {
	var jsonResult Results
	resp, err := http.Get(Url)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error getting all pokemon")
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&jsonResult)
	return c.Status(200).JSON(jsonResult)
}

func SelectedPokemon(c *fiber.Ctx) error {
	type SinglePokemon struct {
		Name string `json:name`
	}
	type Name struct {
		Name string `json:name`
	}
	type Species struct {
		Species Name `json:species`
	}
	// var jsonResult SinglePokemon
	var picture Species
	name := c.Params("name")
	resp, err := http.Get(Url + name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendStatus(404)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(body)
	// json.NewDecoder(resp.Body).Decode(&jsonResult)
	json.NewDecoder(resp.Body).Decode(&picture.Species.Name)

	return c.JSON(picture)
}
