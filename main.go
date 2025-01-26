package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	// If you want to use another engine,
	// just replace with following:
	// Create a new engine with django
	// engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title":       "Go Fiber Template Example",
			"Description": "An example template",
			"Greeting":    "Hello, world!",
		})
	})

	log.Fatal(app.Listen(":3000"))
}