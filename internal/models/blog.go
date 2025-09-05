package models

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null;size:255"`
	Content   string         `json:"content" gorm:"not null;type:text"`
	Slug      string         `json:"slug" gorm:"unique;not null;size:255"`
	Published bool           `json:"published" gorm:"default:false"`
	ImageURL  string         `json:"image_url" gorm:"size:255"`
	ImageID   string         `json:"image_id" gorm:"size:255"`
	UserID    uint           `json:"userId" gorm:"not null"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type BlogResponse struct {
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	Slug      string       `json:"slug"`
	Published bool         `json:"published"`
	UserID    uint         `json:"userId"`
	User      UserResponse `json:"user"`
	ImageURL  string       `json:"image_url"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type CreateBlogRequest struct {
	Title     string          `json:"title" validate:"required,min=1,max=255"`
	Content   string          `json:"content" validate:"required,min=1"`
	Published bool            `json:"published"`
	UserID    uint            `json:"user_id"`
	Image     *multipart.FileHeader `json:"image" form:"image"`
}

type UpdateBlogRequest struct {
	Title     string                `json:"title" validate:"required,min=1,max=255"`
	Content   string                `json:"content" validate:"required,min=1"`
	Published *bool                 `json:"published"`
	Image     *multipart.FileHeader `json:"image" form:"image"`
}

type BlogQueryParams struct {
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
	Search    string `query:"search"`
	Published *bool  `query:"published"`
	UserID    uint   `query:"user_id"`
}

func (u *Blog) ToResponse() BlogResponse {
	return BlogResponse{
		ID:        u.ID,
		Title:     u.Title,
		Content:   u.Content,
		Slug:      u.Slug,
		Published: u.Published,
		UserID:    u.UserID,
		User:      u.User.ToResponse(),
		ImageURL:  u.ImageURL,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
