package repositories

import (
	"http-api/dtos"
	"http-api/entities"

	"gorm.io/gorm"
)

type IAnswerRepository interface {
	FindAll() ([]entities.Answer, error)
	FindById(id int64) (entities.Answer, error)
	Create(answer entities.Answer) error
	Update(id int64, updateInput dtos.UpdateAnswerInput) error
	Delete(id int64) error
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

func (rep *answerRepository) FindById(id int64) (entities.Answer, error) {
	var result entities.Answer
	err := rep.db.First(&result, entities.Answer{
		ID: id,
	}).Error
	return result, err
}

func (rep *answerRepository) Update(id int64, updateInput dtos.UpdateAnswerInput) error {
	// Update the value according to the id and the input fields
	err := rep.db.Model(&entities.Answer{}).Where("ID = ?", id).Omit("ID").Updates(updateInput).Error
	return err
}

func (rep *answerRepository) Delete(id int64) error {
	// SOFT DELETE the value according to the id
	err := rep.db.Delete(&entities.Answer{
		ID: id,
	}).Error
	return err
}
