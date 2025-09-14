package services

import (
	"HR-User-Service/internal/models"
	"HR-User-Service/internal/repository"
)

type RoleService struct {
	repo *repository.RoleRepository
}

func NewRoleService(repo *repository.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) CreateRole(req models.CreateRoleRequest) (*models.Role, error) {
	role := &models.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.repo.Create(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}
