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
	return role, err
}

func (s *RoleService) GetRoleByID(id string) (*models.Role, error) {
	return s.repo.GetByID(id)
}

func (s *RoleService) GetRoleByName(name string) (*models.Role, error) {
	return s.repo.GetByName(name)
}

func (s *RoleService) ExistsByID(id string) (bool, error) {
	return s.repo.ExistsByID(id)
}

func (s *RoleService) ExistsByName(name string) (bool, error) {
	return s.repo.ExistsByName(name)
}
