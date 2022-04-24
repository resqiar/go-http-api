package repositories

import (
	"http-api/entities"

	"gorm.io/gorm"
)

type IAnswerRepository interface {
	FindAll() ([]entities.Answer, error)
	Create(answer entities.Answer) error
}

type answerRepository struct {
	db *gorm.DB
}

func AnswerRepository(db *gorm.DB) *answerRepository {
	return &answerRepository{db}
}

func (rep *answerRepository) FindAll() ([]entities.Answer, error) {
	var result []entities.Answer
	err := rep.db.Find(&result).Error
	return result, err
}

func (rep *answerRepository) Create(answer entities.Answer) error {
	err := rep.db.Create(&answer).Save(&answer).Error
	return err
}
