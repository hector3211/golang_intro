package controllers

import (
	// "encoding/json"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const Url = "https://pokeapi.co/api/v2/pokemon/"

func GetAllPokemons(c *fiber.Ctx) error {
	type Pokemons struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	type Results struct {
		Results []Pokemons `json:"results"`
	}
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
	// type SelectedPokemon struct {
	// 	Name string `json:"name"`
	// }

	var pokemon map[string]interface{}
	name := c.Params("name")
	resp, err := http.Get(Url + name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendStatus(404)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&pokemon)
	return c.JSON(pokemon)
}

func PokemonPic(c *fiber.Ctx) error {
	type Front_Default struct {
		Front_Default string `json:"front_default"`
	}
	type Dream_World struct {
		Dream_World Front_Default `json:"dream_world"`
	}
	type Other struct {
		Other Dream_World `json:"other"`
	}
	type Sprites struct {
		Sprites Other `json:"sprites"`
	}

	var pokemonPic Sprites
	name := c.Params("name")
	resp, err := http.Get(Url + name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendStatus(404)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&pokemonPic)
	return c.JSON(pokemonPic.Sprites.Other.Dream_World)
}
