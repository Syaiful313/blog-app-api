package controllers

import (
	"strconv"

	"go-fiber-boilerplate/config"
	"go-fiber-boilerplate/internal/models"
	"go-fiber-boilerplate/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

func (h *BlogController) GetBlogs(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	blogs, total, err := h.blogService.GetBlogs(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch blogs"})
	}

	var responses []models.BlogResponse
	for _, blog := range blogs {
		responses = append(responses, blog.ToResponse())
	}

	return c.JSON(fiber.Map{
		"data": responses,
		"pagination": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func (h *BlogController) GetBlog(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid blog ID"})
	}

	blog, err := h.blogService.GetBlog(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(fiber.Map{"data": blog.ToResponse()})
}
