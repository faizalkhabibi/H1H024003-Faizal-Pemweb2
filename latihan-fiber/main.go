package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/mahasiswa", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"nim":   "H1H024003",
			"nama":  "Muhammad Faizal Khabibi",
			"prodi": "Teknik Komputer",
		})
	})

	app.Listen(":3000")
}