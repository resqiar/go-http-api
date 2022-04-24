package repositories

import (
	"http-api/entities"

	"gorm.io/gorm"
)

type IQuestionRepository interface {
	FindAll() ([]entities.Question, error)
	Create(question entities.Question) error
}

type questionRepository struct {
	db *gorm.DB
}

func QuestionRepository(db *gorm.DB) *questionRepository {
	return &questionRepository{db}
}

func (rep *questionRepository) FindAll() ([]entities.Question, error) {
	var result []entities.Question
	err := rep.db.Find(&result).Error
	return result, err
}

func (rep *questionRepository) Create(question entities.Question) error {
	err := rep.db.Create(&question).Save(&question).Error
	return err
}
