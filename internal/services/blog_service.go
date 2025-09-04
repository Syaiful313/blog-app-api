package services

import (
	"errors"
	"go-fiber-boilerplate/config"
	"go-fiber-boilerplate/database"
	"go-fiber-boilerplate/internal/models"

	"github.com/gosimple/slug"
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
