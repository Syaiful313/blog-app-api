package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null;size:255"`
	Content   string         `json:"content" gorm:"not null;type:text"`
	Slug      string         `json:"slug" gorm:"unique;not null;size:255"`
	Published bool           `json:"published" gorm:"default:false"`
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
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type CreateBlogRequest struct {
	Title     string `json:"title" validate:"required,min=1,max=255"`
	Content   string `json:"content" validate:"required,min=1"`
	Published bool   `json:"published"`
	Slug      string `json:"slug" validate:"required,min=1,max=255"`
	UserID    uint   `json:"user_id"`
}

type UpdateBlogRequest struct {
	Title     string `json:"title" validate:"required,min=1,max=255"`
	Content   string `json:"content" validate:"required,min=1"`
	Published bool   `json:"published"`
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
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
