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

func NewBlogController(cfg *config.Config) (*BlogController, error) {
	blogService, err := services.NewBlogService(cfg)
	if err != nil {
		return nil, err
	}
	return &BlogController{
		blogService: blogService,
	}, nil
}

func (h *BlogController) CreateBlog(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	title := c.FormValue("title")
	content := c.FormValue("content")
	published := c.FormValue("published") == "true"

	if title == "" || content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and content are required",
		})
	}

	req := models.CreateBlogRequest{
		Title:     title,
		Content:   content,
		Published: published,
	}

	file, err := c.FormFile("image")
	if err == nil {
		req.Image = file
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
func (h *BlogController) GetBlogById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid blog ID"})
	}

	blog, err := h.blogService.GetBlogById(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(fiber.Map{"data": blog.ToResponse()})
}
func (h *BlogController) UpdateBlog(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid blog ID"})
	}

	title := c.FormValue("title")
	content := c.FormValue("content")
	publishedStr := c.FormValue("published")

	req := models.UpdateBlogRequest{
		Title:   title,
		Content: content,
	}

	if publishedStr != "" {
		published := publishedStr == "true"
		req.Published = &published
	}

	file, err := c.FormFile("image")
	if err == nil {
		req.Image = file
	}

	blog, err := h.blogService.UpdateBlog(uint(id), userID, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Blog updated successfully",
		"data":    blog.ToResponse(),
	})
}
