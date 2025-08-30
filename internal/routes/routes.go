package routes

import (
    "go-fiber-boilerplate/config"
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config) {
    // Health check
    app.Get("/api/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status":  "ok", 
            "message": "Server is running",
        })
    })

    // API routes
    api := app.Group("/")

    // Setup route groups
    SetupAuthRoutes(api, cfg)
	SetupSampleRoutes(api, cfg)
}