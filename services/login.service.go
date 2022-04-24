package services

import (
	"http-api/dtos"
	"http-api/entities"
	"http-api/repositories"

	"golang.org/x/crypto/bcrypt"
)

type ILoginService interface {
	Login(email string, password string) (int64, bool)
	HandleRegister(userInput dtos.RegisterInput) error
}

type loginService struct {
	service    IUserService
	repository repositories.IUserRepository
}

func LoginService(s IUserService, r repositories.IUserRepository) *loginService {
	return &loginService{s, r}
}

func (s *loginService) Login(email string, password string) (int64, bool) {
	// Call user service to get the user data by email
	u, err := s.service.FindByEmail(email)
	if err != nil {
		return 0, false
	}

	// Compare the current user password hash with the password input
	// if not match, it will return error.
	hashErr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if hashErr != nil {
		return 0, false
	}
	return u.ID, true
}

func (s *loginService) HandleRegister(userInput dtos.RegisterInput) error {
	// Hash the user's password input
	// This is the best practice to save the password into the database,
	// as if the attacker managed to get access into the database, they still
	// would not know what the actual password is (hashed)
	pwHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Bind input into user obj
	u := entities.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: string(pwHash),
		Bio:      userInput.Bio,
	}

	// Call the repository to create a user
	createErr := s.repository.Create(u)

	// return error (if any)
	return createErr
}
