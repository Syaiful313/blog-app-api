package services

import (
	"errors"
	"go-fiber-boilerplate/config"
	"go-fiber-boilerplate/database"
	"go-fiber-boilerplate/internal/models"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type BlogService struct {
	cfg *config.Config
}

func NewBlogService(cfg *config.Config) *BlogService {
	return &BlogService{cfg: cfg}
}

func (s *BlogService) CreateBlog(userID uint, req models.CreateBlogRequest) (*models.Blog, error) {
	generatedSlug := slug.Make(req.Title)

	var existingBlog models.Blog
	if err := database.GetDB().Where("slug = ?", generatedSlug).First(&existingBlog).Error; err == nil {
		return nil, errors.New("slug already exists")
	}

	blog := models.Blog{
		Title:     req.Title,
		Content:   req.Content,
		UserID:    userID,
		Slug:      generatedSlug,
		Published: req.Published,
	}

	if err := database.GetDB().Create(&blog).Error; err != nil {
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

func (s *BlogService) GetBlog(id uint) (*models.Blog, error) {
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
