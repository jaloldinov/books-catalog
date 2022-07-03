package models

import "time"

type Author struct {
	ID        string    `json:"id" db:"id" binding:"required" example:"uudi1234"`
	Firstname string    `json:"firstname" db:"firstname" binding:"required" example:"John"`
	Lastname  string    `json:"lastname" db:"lastname" binding:"required" example:"Doe"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateAuthor struct {
	Firstname string `json:"firstname" db:"firstname" binding:"required" example:"John"`
	Lastname  string `json:"lastname" db:"lastname" binding:"required" example:"Doe"`
}

type UpdateAuthor struct {
	Firstname string    `json:"firstname" db:"firstname" example:"Updated John"`
	Lastname  string    `json:"lastname" db:"lastname" example:"Updated Doe"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
