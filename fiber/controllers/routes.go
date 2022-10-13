package controllers

import (
	// "encoding/json"
	"encoding/json"
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
	type Front_Default struct {
		Front_Default string `json:front_default`
	}
	type Dream_World struct {
		Dream_World Front_Default `json:dream_world`
	}
	type Other struct {
		Other Dream_World `json:other`
	}
	type Sprites struct {
		Sprites Other `json:sprites`
	}

	// var jsonResult SinglePokemon
	var pokemonPic Sprites
	name := c.Params("name")
	resp, err := http.Get(Url + name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendStatus(404)
	}
	defer resp.Body.Close()
	// json.NewDecoder(resp.Body).Decode(&jsonResult)
	json.NewDecoder(resp.Body).Decode(&pokemonPic)
	return c.JSON(pokemonPic)
}
