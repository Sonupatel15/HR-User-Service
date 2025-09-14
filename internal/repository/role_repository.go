package repository

import (
	"HR-User-Service/internal/database"
	"HR-User-Service/internal/models"
	"time"
)

// RoleRepository handles DB operations for roles
type RoleRepository struct {
	db database.Service
}

func NewRoleRepository(db database.Service) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Create(role *models.Role) error {
	query := `
		INSERT INTO roles (name, description, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	// Get the underlying *sql.DB from the service
	db := r.db.GetDB()

	err := db.QueryRow(
		query,
		role.Name,
		role.Description,
		time.Now(),
	).Scan(&role.ID, &role.CreatedAt)

	return err
}
