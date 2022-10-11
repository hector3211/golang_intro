package main

import (
	"fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/pokemons", controllers.GetAllPokemons)
	app.Get("/pokemons/:name", controllers.SelectedPokemon)
	app.Listen(":3000")
}
