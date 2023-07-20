package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/ilmsg/fiberoptic/fiber-template-blog/internal/router"
)

func main() {
	store := session.New()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.SetupRouter(app, store)

	app.Listen(":3001")
}
