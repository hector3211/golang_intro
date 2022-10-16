package main

import (
	"fiber/controllers"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public/test.html")

	app.Get("/pokemons", controllers.GetAllPokemons)
	app.Get("/pokemons/:name", controllers.SelectedPokemon)
	app.Get("pokemons/:name/pic", controllers.PokemonPic)

	// 404
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// app.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"john":  "doe",
	// 		"admin": "123456",
	// 	},
	// }))
	app.Listen(":3000")
}
