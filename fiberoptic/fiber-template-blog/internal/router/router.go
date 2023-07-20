package router

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App) {
	app.Get("/", GetIndex)
	app.Get("/layout", GetLayout)

	app.Get("/pages", GetPageList)
	app.Get("/pages/:id", GetPageDetail)

	app.Get("/posts", GetPostList)
	app.Get("/posts/:id", GetPostDetail)
}
