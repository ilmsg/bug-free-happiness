package middleware

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ilmsg/cache-api/models"
	"github.com/patrickmn/go-cache"
)

func CacheMiddleware(cache *cache.Cache) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Method() != "GET" {
			// Only cache GET requests
			return ctx.Next()
		}

		cacheKey := ctx.Path() + "?" + ctx.Params("id") // Generate a cache key from the request path and query parameters

		// Check if the response is already in the cache
		if cached, found := cache.Get(cacheKey); found {
			ctx.Response().Header.Set("Cache-Status", "HIT")
			return ctx.JSON(cached)
		}

		ctx.Set("Cache-Status", "MISS")
		err := ctx.Next()
		if err != nil {
			return err
		}

		var data models.Post
		cacheKey = ctx.Path() + "?" + ctx.Params("id")

		body := ctx.Response().Body()
		if err := json.Unmarshal(body, &data); err != nil {
			return ctx.JSON(fiber.Map{"error": err.Error()})
		}

		// Cache the response for 10 minute
		cache.Set(cacheKey, data, 10*time.Minute)

		return nil
	}
}
