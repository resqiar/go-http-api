package repositories

import (
	"http-api/dtos"
	"http-api/entities"

	"gorm.io/gorm"
)

type IQuestionRepository interface {
	FindAll() ([]entities.Question, error)
	FindById(id int64) (entities.Question, error)
	Create(question entities.Question) error
	Update(id int64, updateInput dtos.UpdateQuestionInput) error
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

func (rep *questionRepository) FindById(id int64) (entities.Question, error) {
	var result entities.Question
	err := rep.db.First(&result, entities.Question{
		ID: id,
	}).Error
	return result, err
}

func (rep *questionRepository) Update(id int64, updateInput dtos.UpdateQuestionInput) error {
	// Update the value according to the id and the input fields
	err := rep.db.Model(&entities.Question{}).Where("ID = ?", id).Omit("ID").Updates(updateInput).Error
	return err
}
