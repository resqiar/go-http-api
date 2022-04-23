package users

type IUserService interface {
	FindAll() ([]User, error)
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (SafeUser, error)
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

func (s *service) FindByEmail(email string) (User, error) {
	// Call user service to find user based on email
	result, err := s.repo.FindByEmail(email)
	return result, err
}

func (s *service) FindByUsername(username string) (SafeUser, error) {
	// Call user service to find user based on username
	result, err := s.repo.FindByUsername(username)
	return result, err
}
