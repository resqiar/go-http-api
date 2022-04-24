package repositories

import (
	"http-api/entities"

	"gorm.io/gorm"
)

type IQuestionRepository interface {
	FindAll() ([]entities.Question, error)
	Create(question entities.Question) error
}

type questionrepository struct {
	db *gorm.DB
}

func QuestionRepository(db *gorm.DB) *questionrepository {
	return &questionrepository{db}
}

func (rep *questionrepository) FindAll() ([]entities.Question, error) {
	var result []entities.Question
	err := rep.db.Find(&result).Error
	return result, err
}

func (rep *questionrepository) Create(question entities.Question) error {
	err := rep.db.Create(&question).Save(&question).Error
	return err
}
