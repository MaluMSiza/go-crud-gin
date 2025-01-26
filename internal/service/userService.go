package service

import (
	"github.com/MaLuMSiza/go-crud-gin/internal/models"
	"github.com/MaLuMSiza/go-crud-gin/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userServiceImpl) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}
