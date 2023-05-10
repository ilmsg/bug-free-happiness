package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ilmsg/cache-api/models"
)

func GetHello(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}

func GetPostsById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Fatal("Invalid ID")
	}

	apiUrl := "https://jsonplaceholder.typicode.com/posts/" + id

	// Fetch the post data from the API
	resp, err := http.Get(apiUrl)
	if err != nil {
		return ctx.JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ctx.JSON(fiber.Map{"error": err.Error()})
	}

	var data models.Post
	if err := json.Unmarshal(body, &data); err != nil {
		return ctx.JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(data)
}
