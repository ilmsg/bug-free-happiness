package router

import "github.com/gofiber/fiber/v2"

func GetPageList(ctx *fiber.Ctx) error {
	return ctx.Render("pages/list", fiber.Map{
		"title":       "Title Page",
		"description": "Description Page",
	}, "layouts/main")
}

func GetPageDetail(ctx *fiber.Ctx) error {
	return ctx.Render("pages/detail", fiber.Map{
		"title":       "Title Page",
		"description": "Description Page",
	}, "layouts/main")
}
