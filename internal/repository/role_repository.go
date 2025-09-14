package repository

import (
	"HR-User-Service/internal/database"
	"HR-User-Service/internal/models"
)

type RoleRepository struct {
	db database.Service
}

func NewRoleRepository(db database.Service) *RoleRepository {
	return &RoleRepository{db: db}
}

// Create a new role
func (r *RoleRepository) Create(role *models.Role) error {
	query := `
		INSERT INTO roles (name, description, created_at)
		VALUES ($1, $2, NOW())
		RETURNING id, created_at
	`

	// Get the underlying *sql.DB from the service
	db := r.db.GetDB()

	return db.QueryRow(
		query,
		role.Name,
		role.Description,
	).Scan(&role.ID, &role.CreatedAt)
}

// Get role by ID
func (r *RoleRepository) GetByID(id string) (*models.Role, error) {
	role := &models.Role{}
	query := `SELECT id, name, description, created_at FROM roles WHERE id = $1`

	db := r.db.GetDB()
	err := db.QueryRow(query, id).Scan(
		&role.ID, &role.Name, &role.Description, &role.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// Get role by name
func (r *RoleRepository) GetByName(name string) (*models.Role, error) {
	role := &models.Role{}
	query := `SELECT id, name, description, created_at FROM roles WHERE name = $1`

	db := r.db.GetDB()
	err := db.QueryRow(query, name).Scan(
		&role.ID, &role.Name, &role.Description, &role.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// Check if role exists by ID
func (r *RoleRepository) ExistsByID(id string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM roles WHERE id = $1)`

	db := r.db.GetDB()
	err := db.QueryRow(query, id).Scan(&exists)
	return exists, err
}

// Check if role exists by Name
func (r *RoleRepository) ExistsByName(name string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM roles WHERE name = $1)`

	db := r.db.GetDB()
	err := db.QueryRow(query, name).Scan(&exists)
	return exists, err
}
