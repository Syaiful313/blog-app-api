package routes

import (
	"go-fiber-boilerplate/config"
	"go-fiber-boilerplate/internal/controllers"
	"go-fiber-boilerplate/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupBlogRouter(api fiber.Router, cfg *config.Config) {
	blogController := controllers.NewBlogController(cfg)

	blogs := api.Group("/blogs")

	blogs.Post("/", middleware.AuthMiddleware(cfg), blogController.CreateBlog)
	blogs.Get("/", blogController.GetBlogs)
	blogs.Get("/:id", blogController.GetBlog)
}
