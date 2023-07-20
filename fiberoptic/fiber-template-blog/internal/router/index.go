package router

import "github.com/gofiber/fiber/v2"

func GetIndex(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{
		"title":       "Title Page",
		"description": "Description Page",
	})
}

func GetLayout(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{
		"title":       "Title Page",
		"description": "Description Page",
	}, "layouts/main")
}
