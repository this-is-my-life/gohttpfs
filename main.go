package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pmh-only/gohttpfs/configloader"
	"github.com/pmh-only/gohttpfs/flagsolver"
	"github.com/pmh-only/gohttpfs/middlewares"
)

var flags flagsolver.Flags

func init() {
	flags = flagsolver.SolveFlag()
}

func main() {
	app := fiber.New(fiber.Config{AppName: "gohttpfs", Prefork: true})
	config := configloader.LoadConfig(*flags.ConfigFilePath, flags.Configuration)

	app.Use(logger.New())
	app.Use(*config.ServePrefix, middlewares.ExtraHeaders(config))

	app.Static(*config.ServePrefix, *config.StoragePath, fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Index:         *config.DefaultDocument,
		CacheDuration: time.Duration(*config.CacheDuration) * time.Second,
		Next: func(c *fiber.Ctx) bool {
			return c.Method() != "GET" || len(c.Context().URI().QueryString()) > 0 || len(c.Get("X-Client")) > 0
		},
	})

	app.Use(*config.ServePrefix, middlewares.ArchiveApi(config))
	app.Use(*config.ServePrefix, middlewares.ListApi(config))
	app.Listen(*config.Listen)
}
