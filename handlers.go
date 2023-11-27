package main

import (
	"github.com/gofiber/fiber/v2"
	"sync"
)

var mut sync.Mutex

var Items = &[]string{"apple", "orange", "pineapple"}

func home(c *fiber.Ctx) error {
	return c.Render("index", nil)

}

func list(c *fiber.Ctx) error {
	return c.Render("list", *Items)

}

func add(c *fiber.Ctx) error {
	return c.Render("add", *Items)

}

func addItem(c *fiber.Ctx) error {

	fn := c.FormValue("fruit")
	if fn != "" {
		mut.Lock()
		defer mut.Unlock()
		*Items = append(*Items, fn)
		return c.Render("add", Items)
	}

	return c.Render("empty", "Error: Empty data!")
}
