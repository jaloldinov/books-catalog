package models

import "time"

type BookCategory struct {
	ID           string    `json:"id" db:"id" example:"123abd"`
	CategoryName string    `json:"category_name" db:"category_name" binding:"required" example:"fiction"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type CreateBookCategory struct {
	CategoryName string `json:"category_name" db:"category_name" binding:"required" example:"fiction"`
}

type UpdateBookCategory struct {
	CategoryName string    `json:"category_name" db:"category_name" binding:"required" example:"fiction"`
}
