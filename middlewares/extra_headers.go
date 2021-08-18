package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pmh-only/gohttpfs/configloader"
)

func ExtraHeaders(config configloader.Configuration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return err
		}

		for k, v := range config.ExtraHeaders {
			c.Set(k, v)
		}

		return nil
	}
}
