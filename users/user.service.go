package users

import "golang.org/x/crypto/bcrypt"

type IUserService interface {
	FindAll() ([]User, error)
	Create(userInput UserInput) (int, error)
}

type service struct {
	repo IUserRepository
}

func UserService(repo IUserRepository) *service {
	return &service{repo}
}

func (s *service) FindAll() ([]User, error) {
	result, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *service) Create(userInput UserInput) (int, error) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return 400, err
	}
	u := User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: string(pwHash),
		Bio:      userInput.Bio,
	}
	result, err := s.repo.Create(u)
	return result, err
}
