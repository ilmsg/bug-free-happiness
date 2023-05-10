package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ilmsg/cache-api/middleware"
	"github.com/ilmsg/cache-api/routes"
	"github.com/patrickmn/go-cache"
)

func main() {
	app := fiber.New()

	cache := cache.New(10*time.Minute, 20*time.Minute) // setting default expiration time and clearance time.

	app.Get("/", routes.GetHello)
	app.Get("/posts/:id", middleware.CacheMiddleware(cache), routes.GetPostsById)

	app.Listen(":4000")
}
