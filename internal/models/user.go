package models

import "time"

type User struct {
	ID             string    `db:"id" json:"id"`
	Name           string    `db:"name" json:"name"`
	PrimaryEmail   string    `db:"primary_email" json:"primary_email"`
	SecondaryEmail string    `db:"secondary_email,omitempty" json:"secondary_email,omitempty"`
	MobileNumber   string    `db:"mobile_number" json:"mobile_number"`
	Password       string    `db:"password" json:"-"` // don’t expose in JSON
	RoleID         string    `db:"role_id" json:"role_id"`
	IsActive       bool      `db:"is_active" json:"is_active"`
	LastLogin      time.Time `db:"last_login,omitempty" json:"last_login,omitempty"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
