package handlers

import "github.com/gofiber/fiber"

func welcome(c *fiber.Ctx) error {
	return c.Render("welcome", nil, "layouts/main")
}