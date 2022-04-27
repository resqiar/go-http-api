package repositories

import (
	"http-api/entities"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAll() ([]entities.User, error)
	FindByEmail(email string) (entities.User, error)
	FindByUsername(username string) (entities.SafeUser, error)
	FindById(id int64) (entities.SafeUser, error)
	Create(user entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (rep *userRepository) FindAll() ([]entities.User, error) {
	var result []entities.User

	// Find all users with its relations ("Questions")
	err := rep.db.Preload("Questions").Find(&result).Error
	return result, err
}

func (rep *userRepository) FindByEmail(email string) (entities.User, error) {
	var result entities.User

	// Find the first match user by email.
	err := rep.db.Where("email = ?", email).First(&result).Error
	return result, err
}

func (rep *userRepository) FindByUsername(username string) (entities.SafeUser, error) {
	var result entities.SafeUser

	// Find the first match user by username.
	// Omit the password as it is used as a public endpoint.
	// Smart query := https://gorm.io/docs/advanced_query.html#Smart-Select-Fields
	err := rep.db.Model(&entities.User{}).First(&result, entities.User{
		Username: username,
	}).Error

	return result, err
}

func (rep *userRepository) FindById(id int64) (entities.SafeUser, error) {
	var result entities.SafeUser

	// Find the first match user by id.
	// Omit the password as it is used as a public endpoint.
	// Smart query := https://gorm.io/docs/advanced_query.html#Smart-Select-Fields
	err := rep.db.Model(&entities.User{}).Where("ID = ?", id).First(&result).Error

	return result, err
}

func (repo *userRepository) Create(user entities.User) error {
	// Create and save to DB
	err := repo.db.Create(&user).Save(&user).Error
	return err
}
