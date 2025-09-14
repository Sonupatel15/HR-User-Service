package models

import "time"

type Role struct {
	ID          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description,omitempty" json:"description,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

// DTO for request payload
type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
