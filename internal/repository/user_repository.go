package repository

import (
	"HR-User-Service/internal/database"
	"HR-User-Service/internal/models"
	"database/sql"
)

type UserRepository struct {
	db database.Service
}

func NewUserRepository(db database.Service) *UserRepository {
	return &UserRepository{db: db}
}

// Create a new user
func (u *UserRepository) Create(user *models.User) error {
	query := `
        INSERT INTO users (
            name, primary_email, secondary_email, mobile_number, 
            password, role_id, is_active, last_login
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, created_at, updated_at
    `

	db := u.db.GetDB()

	// Handle NULL values
	secondaryEmail := sql.NullString{String: user.SecondaryEmail, Valid: user.SecondaryEmail != ""}
	mobileNumber := sql.NullString{String: user.MobileNumber, Valid: user.MobileNumber != ""}
	lastLogin := sql.NullTime{Time: user.LastLogin, Valid: !user.LastLogin.IsZero()}

	return db.QueryRow(
		query,
		user.Name,
		user.PrimaryEmail,
		secondaryEmail,
		mobileNumber,
		user.Password,
		user.RoleID,
		user.IsActive,
		lastLogin,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

// Get user by ID - FIXED (added missing fields)
func (u *UserRepository) GetByID(id string) (*models.User, error) {
	user := &models.User{}
	query := `
        SELECT id, name, primary_email, secondary_email, mobile_number, 
               password, role_id, is_active, last_login, created_at, updated_at 
        FROM users WHERE id = $1
    `

	db := u.db.GetDB()

	// Handle NULL values
	var secondaryEmail, mobileNumber sql.NullString
	var lastLogin sql.NullTime

	err := db.QueryRow(query, id).Scan(
		&user.ID, &user.Name, &user.PrimaryEmail, &secondaryEmail, &mobileNumber,
		&user.Password, &user.RoleID, &user.IsActive, &lastLogin,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Convert NULL values to empty strings/zero time
	user.SecondaryEmail = secondaryEmail.String
	user.MobileNumber = mobileNumber.String
	if lastLogin.Valid {
		user.LastLogin = lastLogin.Time
	}

	return user, nil
}

// Get user by name - FIXED (added missing fields)
func (u *UserRepository) GetByName(name string) (*models.User, error) {
	user := &models.User{}
	query := `
        SELECT id, name, primary_email, secondary_email, mobile_number, 
               password, role_id, is_active, last_login, created_at, updated_at 
        FROM users WHERE name = $1
    `

	db := u.db.GetDB()

	var secondaryEmail, mobileNumber sql.NullString
	var lastLogin sql.NullTime

	err := db.QueryRow(query, name).Scan(
		&user.ID, &user.Name, &user.PrimaryEmail, &secondaryEmail, &mobileNumber,
		&user.Password, &user.RoleID, &user.IsActive, &lastLogin,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	user.SecondaryEmail = secondaryEmail.String
	user.MobileNumber = mobileNumber.String
	if lastLogin.Valid {
		user.LastLogin = lastLogin.Time
	}

	return user, nil
}

// Get user by primary email - FIXED (typo in column name)
func (u *UserRepository) GetByPrimaryEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
        SELECT id, name, primary_email, secondary_email, mobile_number, 
               password, role_id, is_active, last_login, created_at, updated_at 
        FROM users WHERE primary_email = $1
    `

	db := u.db.GetDB()

	var secondaryEmail, mobileNumber sql.NullString
	var lastLogin sql.NullTime

	err := db.QueryRow(query, email).Scan(
		&user.ID, &user.Name, &user.PrimaryEmail, &secondaryEmail, &mobileNumber,
		&user.Password, &user.RoleID, &user.IsActive, &lastLogin,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	user.SecondaryEmail = secondaryEmail.String
	user.MobileNumber = mobileNumber.String
	if lastLogin.Valid {
		user.LastLogin = lastLogin.Time
	}

	return user, nil
}

// Get user by secondary email - FIXED
func (u *UserRepository) GetBySecondaryEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
        SELECT id, name, primary_email, secondary_email, mobile_number, 
               password, role_id, is_active, last_login, created_at, updated_at 
        FROM users WHERE secondary_email = $1
    `

	db := u.db.GetDB()

	var secondaryEmail, mobileNumber sql.NullString
	var lastLogin sql.NullTime

	err := db.QueryRow(query, email).Scan(
		&user.ID, &user.Name, &user.PrimaryEmail, &secondaryEmail, &mobileNumber,
		&user.Password, &user.RoleID, &user.IsActive, &lastLogin,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	user.SecondaryEmail = secondaryEmail.String
	user.MobileNumber = mobileNumber.String
	if lastLogin.Valid {
		user.LastLogin = lastLogin.Time
	}

	return user, nil
}

// Check if user exists by ID - FIXED (typo in EXISTS)
func (u *UserRepository) ExistsByID(id string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
	db := u.db.GetDB()
	err := db.QueryRow(query, id).Scan(&exists)
	return exists, err
}

// Check if user exists by Name - FIXED
func (u *UserRepository) ExistsByName(name string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE name = $1)`
	db := u.db.GetDB()
	err := db.QueryRow(query, name).Scan(&exists)
	return exists, err
}

// Check if user exists by PrimaryEmail - FIXED (wrong receiver type)
func (u *UserRepository) ExistsByPrimaryEmail(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE primary_email = $1)`
	db := u.db.GetDB()
	err := db.QueryRow(query, email).Scan(&exists)
	return exists, err
}

// Check if user exists by SecondaryEmail - FIXED (wrong receiver type)
func (u *UserRepository) ExistsBySecondaryEmail(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE secondary_email = $1)`
	db := u.db.GetDB()
	err := db.QueryRow(query, email).Scan(&exists)
	return exists, err
}
