package users

import "gorm.io/gorm"

type IUserRepository interface {
	FindAll() ([]User, error)
	Create(user User) (int, error)
}

type repository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (rep *repository) FindAll() ([]User, error) {
	var result []User
	err := rep.db.Find(&result).Error
	return result, err
}

func (repo *repository) Create(user User) (int, error) {
	err := repo.db.Create(&user).Error
	return 200, err
}
