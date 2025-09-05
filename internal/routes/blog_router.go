package routes

import (
	"go-fiber-boilerplate/config"
	"go-fiber-boilerplate/internal/controllers"
	"go-fiber-boilerplate/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupBlogRouter(api fiber.Router, cfg *config.Config) {
	blogController, err := controllers.NewBlogController(cfg)
	if err != nil {
		panic(err)
	}

	blogs := api.Group("/blogs")

	blogs.Post("/",
		middlewares.AuthMiddleware(cfg),
		middlewares.NewUploaderMiddleware().ImageUpload(2, []string{"image/jpeg", "image/png"}),
		blogController.CreateBlog,
	)
	blogs.Get("/", blogController.GetBlogs)
	blogs.Get("/:id", blogController.GetBlogById)
	blogs.Patch("/:id",
		middlewares.AuthMiddleware(cfg),
		middlewares.NewUploaderMiddleware().ImageUpload(2, []string{"image/jpeg", "image/png"}),
		blogController.UpdateBlog,
	)
}
	