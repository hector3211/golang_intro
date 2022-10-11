package controllers

import "github.com/gofiber/fiber/v2"

func GetValue(c *fiber.Ctx) error {
	return c.SendString("value is:" + c.Params("value"))
}
