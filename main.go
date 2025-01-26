package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	engine := html.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// add static folder
	app.Static(
		"/static", // mount address
		"./web",   // path to the file folder
	)

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
