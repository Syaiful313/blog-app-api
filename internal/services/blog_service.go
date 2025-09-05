package services

import (
	"errors"
	"fmt"
	"go-fiber-boilerplate/config"
	"go-fiber-boilerplate/database"
	"go-fiber-boilerplate/internal/models"
	"go-fiber-boilerplate/utils"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type BlogService struct {
	cfg        *config.Config
	cloudinary *utils.Cloudinary
}

func NewBlogService(cfg *config.Config) (*BlogService, error) {
	cloudinary, err := utils.NewCloudinaryService(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cloudinary: %w", err)
	}
	return &BlogService{
		cfg:        cfg,
		cloudinary: cloudinary,
	}, nil
}

func (s *BlogService) CreateBlog(userID uint, req models.CreateBlogRequest) (*models.Blog, error) {
	var existingBlog models.Blog
	if err := database.GetDB().Where("title = ?", req.Title).First(&existingBlog).Error; err == nil {
		return nil, errors.New("title already exists")
	}

	generatedSlug := slug.Make(req.Title)

	blog := models.Blog{
		Title:     req.Title,
		Content:   req.Content,
		UserID:    userID,
		Slug:      generatedSlug,
		Published: req.Published,
	}

	if req.Image != nil {
		uploadResult, err := s.cloudinary.UploadImage(req.Image, "blog-images")
		if err != nil {
			return nil, fmt.Errorf("failed to upload image: %w", err)
		}
		blog.ImageURL = uploadResult.SecureURL
		blog.ImageID = uploadResult.PublicID
	}

	if err := database.GetDB().Create(&blog).Error; err != nil {
		if blog.ImageID != "" {
			_ = s.cloudinary.DeleteImage(blog.ImageID)
		}
		return nil, err
	}

	database.GetDB().Preload("User").First(&blog, blog.ID)
	return &blog, nil
}

func (s *BlogService) GetBlogs(page, limit int) ([]models.Blog, int64, error) {
	var blogs []models.Blog
	offset := (page - 1) * limit

	if err := database.GetDB().
		Preload("User").
		Offset(offset).
		Limit(limit).
		Find(&blogs).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := database.GetDB().Model(&models.Blog{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return blogs, total, nil
}

func (s *BlogService) GetBlogById(id uint) (*models.Blog, error) {
	var blog models.Blog
	if err := database.GetDB().
		Preload("User").
		First(&blog, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &blog, nil
}

func (s *BlogService) UpdateBlog(id uint, userId uint, req models.UpdateBlogRequest) (*models.Blog, error) {
	var blog models.Blog
	if err := database.GetDB().First(&blog, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if blog.UserID != userId {
		return nil, errors.New("unauthorized to update this blog")
	}

	if req.Title != "" && req.Title != blog.Title {
		generatedSlug := slug.Make(req.Title)
		var existingBlog models.Blog
		if err := database.GetDB().Where("title = ? AND id != ?", req.Title, id).First(&existingBlog).Error; err == nil {
			return nil, errors.New("title already exists")
		}
		blog.Slug = generatedSlug
		blog.Title = req.Title
	}

	if req.Content != "" {
		blog.Content = req.Content
	}

	if req.Published != nil {
		blog.Published = *req.Published
	}

	if req.Image != nil {
		if blog.ImageID != "" {
			err := s.cloudinary.DeleteImage(blog.ImageID)
			if err != nil {
				return nil, fmt.Errorf("failed to delete old image: %w", err)
			}
		}

		uploadResult, err := s.cloudinary.UploadImage(req.Image, "blog-images")
		if err != nil {
			return nil, fmt.Errorf("failed to upload new image: %w", err)
		}
		blog.ImageURL = uploadResult.SecureURL
		blog.ImageID = uploadResult.PublicID
	}

	if err := database.GetDB().Save(&blog).Error; err != nil {
		if req.Image != nil && blog.ImageID != "" {
			_ = s.cloudinary.DeleteImage(blog.ImageID)
		}
		return nil, err
	}

	database.GetDB().Preload("User").First(&blog, blog.ID)
	return &blog, nil
}
