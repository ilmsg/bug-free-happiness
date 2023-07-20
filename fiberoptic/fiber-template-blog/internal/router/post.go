package router

import "github.com/gofiber/fiber/v2"

func GetPostList(ctx *fiber.Ctx) error {
	return ctx.Render("posts/list", fiber.Map{
		"title":       "Title Post",
		"description": "Description Post",
	}, "layouts/main")
}

func GetPostDetail(ctx *fiber.Ctx) error {
	return ctx.Render("posts/detail", fiber.Map{
		"title":       "Title Post",
		"description": "Description Post",
	}, "layouts/main")
}
