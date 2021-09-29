package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"hits/api/prisma/db"
	"hits/api/utils"
	. "hits/api/utils"
	. "hits/api/v1"
	"log"
	"os"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(Response{
			Success: true,
			Message: "Welcome to Hits API!",
		})
	})

	/* V1 */
	v1 := app.Group("/v1")
	v1.Get("/top", utils.RateLimit(50), utils.CacheRoute(), GetTopHits)
	v1.Get("/hits", utils.RateLimit(15), GetHits)
}

func main() {
	_ = godotenv.Load()

	app := fiber.New(fiber.Config{
		CaseSensitive: false,
		StrictRouting: true,
		ServerHeader:  "Hits API",
		AppName:       "Hits API v1.0",
		BodyLimit:     1024 * 1024,
		GETOnly: true,
	})

	app.Use(logger.New(logger.Config{
		Format: "${time} |   ${cyan}${status} ${reset}|   ${latency} | ${ip} on ${cyan}${ua} ${reset}| ${cyan}${method} ${reset}${path} \n",
	}))

	app.Use(recover.New(recover.Config{
		Next:             nil,
		EnableStackTrace: true,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	setupRoutes(app)

	if utils.GetPrisma() == nil {
		utils.SetGlobalDb(db.NewClient())
	}

	if err := utils.GetPrisma().Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := utils.GetPrisma().Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
