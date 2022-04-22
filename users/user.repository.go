package users

import "gorm.io/gorm"

type IUserRepository interface {
	FindAll() ([]User, error)
	FindByEmail(email string) (User, error)
	Create(user User) error
}

type repository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (rep *repository) FindAll() ([]User, error) {
	var result []User

	// Find all users with its relations ("Tasks")
	err := rep.db.Preload("Tasks").Find(&result).Error
	return result, err
}

func (rep *repository) FindByEmail(email string) (User, error) {
	var result User

	// Find the first match user by email.
	err := rep.db.First(&result, User{
		Email: email,
	}).Error
	return result, err
}

func (repo *repository) Create(user User) error {
	// Create and save to DB
	err := repo.db.Create(&user).Save(&user).Error
	return err
}
