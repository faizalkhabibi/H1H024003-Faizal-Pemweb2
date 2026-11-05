package middleware

import "github.com/gofiber/fiber/v3"

func KunciApi(kunci string) fiber.Handler {
	return func(c fiber.Ctx) error {
		apiKey := c.Get("X-Api-Key")

		if kunci == "" || apiKey != kunci {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"API key tidak valid",
			)
		}

		return c.Next()
	}
}
