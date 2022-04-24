package services

import (
	"http-api/entities"
	"http-api/repositories"
)

type IUserService interface {
	FindAll() ([]entities.User, error)
	FindByEmail(email string) (entities.User, error)
	FindByUsername(username string) (entities.SafeUser, error)
}

type userservice struct {
	repo repositories.IUserRepository
}

func UserService(repo repositories.IUserRepository) *userservice {
	return &userservice{repo}
}

func (s *userservice) FindAll() ([]entities.User, error) {
	result, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *userservice) FindByEmail(email string) (entities.User, error) {
	// Call user service to find user based on email
	result, err := s.repo.FindByEmail(email)
	return result, err
}

func (s *userservice) FindByUsername(username string) (entities.SafeUser, error) {
	// Call user service to find user based on username
	result, err := s.repo.FindByUsername(username)
	return result, err
}
