package router

import "github.com/gofiber/fiber/v2"

func GetUserLogin(ctx *fiber.Ctx) error {
	return ctx.Render("users/login", fiber.Map{
		"title":       "Title Post",
		"description": "Description Post",
	}, "layouts/main")
}

func PostUserLogin(ctx *fiber.Ctx) error {
	return ctx.Render("users/login", fiber.Map{
		"title":       "Title Post",
		"description": "Description Post",
	}, "layouts/main")
}

func GetUserRegister(ctx *fiber.Ctx) error {
	return ctx.Render("users/register", fiber.Map{
		"title":       "Title Post",
		"description": "Description Post",
	}, "layouts/main")
}

func PostUserRegister(ctx *fiber.Ctx) error {
	return ctx.Render("users/register", fiber.Map{
		"title":       "Title Post",
		"description": "Description Post",
	}, "layouts/main")
}

func GetUserProfile(ctx *fiber.Ctx) error {
	return ctx.Render("users/register", fiber.Map{
		"title":       "Title Post",
		"description": "Description Post",
	}, "layouts/main")
}

func PostUserProfile(ctx *fiber.Ctx) error {
	return ctx.Render("users/register", fiber.Map{
		"title":       "Title Post",
		"description": "Description Post",
	}, "layouts/main")
}
