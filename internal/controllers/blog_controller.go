package controllers

import (
	"go-fiber-boilerplate/config"
	"go-fiber-boilerplate/internal/models"
	"go-fiber-boilerplate/internal/services"

	"github.com/gofiber/fiber/v2"
)

type BlogController struct {
	blogService *services.BlogService
}

func NewBlogController(cfg *config.Config) *BlogController {
	return &BlogController{
		blogService: services.NewBlogService(cfg),
	}
}

func (h *BlogController) CreateBlog(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req models.CreateBlogRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Title == "" || req.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and content are required",
		})
	}

	blog, err := h.blogService.CreateBlog(userID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Blog created successfully",
		"data":    blog.ToResponse(),
	})
}
