package answers

import "gorm.io/gorm"

type IAnswerRepository interface {
	FindAll() ([]Answer, error)
	Create(answer Answer) error
}

type repository struct {
	db *gorm.DB
}

func AnswerRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (rep *repository) FindAll() ([]Answer, error) {
	var result []Answer
	err := rep.db.Find(&result).Error
	return result, err
}

func (rep *repository) Create(answer Answer) error {
	err := rep.db.Create(&answer).Save(&answer).Error
	return err
}
