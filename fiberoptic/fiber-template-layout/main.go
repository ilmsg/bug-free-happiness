package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"title":       "Title Page",
			"description": "Description Page",
		})
	})

	app.Get("/layout", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"title":       "Title Page",
			"description": "Description Page",
		}, "layouts/main")
	})

	app.Listen(":3001")
}
