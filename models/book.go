package models

import "time"

type Book struct {
	ID         string    `json:"id" db:"id" example:"uuid1234"`
	BookName   string    `json:"book_name" db:"name" binding:"required" example:"book name"`
	AuthorID   string    `json:"author_id" db:"author_id" binding:"required"`
	CategoryID string    `json:"category_id" db:"category_id" binding:"required" example:"uuid1234"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type CreateBook struct {
	CategoryID string `json:"category_id" db:"category_id" binding:"required" example:"uuid1234"`
	AuthorID   string `json:"author_id" db:"author_id" binding:"required" example:"author_id"`
	Name       string `json:"name" db:"name" binding:"required" example:"Start with why"`
}

type UpdateBook struct {
	AuthorID   string    `json:"author_id" db:"author_id" example:"uudi1234"`
	CategoryID string    `json:"category_id" db:"category_id" example:"uudi1234"`
	Name       string    `json:"name" db:"name" example:"Updated book name"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
