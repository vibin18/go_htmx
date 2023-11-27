package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	engine := html.New("templates", ".html")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", home)
	app.Get("/list", list)
	app.Get("/add", add)
	app.Post("/addItem", addItem)
	log.Fatal(app.Listen(":3000"))
}
