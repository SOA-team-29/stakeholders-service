package model

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type BlogPostStatus int

const (
	Draft BlogPostStatus = iota
	Published
	Closed
)

type ImageIDs struct {
	IDs []int
}

type Blog struct {
	ID                uuid.UUID      `json:"id"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	PublishedDateTime *time.Time     `json:"publishedDateTime,omitempty"`
	ImageIDs          []int          `json:"image"`
	Status            BlogPostStatus `json:"status"`
}

func (blog *Blog) BeforeCreate(scope *gorm.DB) error {
	blog.ID = uuid.New()
	return nil
}
