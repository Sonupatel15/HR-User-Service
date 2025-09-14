package services

import (
	"HR-User-Service/internal/models"
	"HR-User-Service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		Name:           req.Name,
		PrimaryEmail:   req.PrimaryEmail,
		SecondaryEmail: req.SecondaryEmail,
		MobileNumber:   req.MobileNumber,
		Password:       req.Password,
		RoleID:         req.RoleID,
	}
	err := s.repo.Create(user)
	return user, err

}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetUserByName(name string) (*models.User, error) {
	return s.repo.GetByName(name)
}

func (s *UserService) GetUserByPrimaryEmail(email string) (*models.User, error) {
	return s.repo.GetByPrimaryEmail(email)
}

func (s *UserService) GetUserBySecondaryEmail(email string) (*models.User, error) {
	return s.repo.GetBySecondaryEmail(email)
}

func (s *UserService) ExistsByID(id string) (bool, error) {
	return s.repo.ExistsByID(id)
}

func (s *UserService) ExistsByName(name string) (bool, error) {
	return s.repo.ExistsByName(name)
}

func (s *UserService) ExistsByPrimaryEmail(email string) (bool, error) {
	return s.repo.ExistsByPrimaryEmail(email)
}

func (s *UserService) ExistsBySecondaryEmail(email string) (bool, error) {
	return s.repo.ExistsBySecondaryEmail(email)
}
